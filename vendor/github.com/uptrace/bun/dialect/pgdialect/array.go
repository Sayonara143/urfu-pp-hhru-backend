package pgdialect

import (
	"database/sql"
	"database/sql/driver"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/uptrace/bun/dialect"
	"github.com/uptrace/bun/internal"
	"github.com/uptrace/bun/schema"
)

type ArrayValue struct {
	v reflect.Value

	append schema.AppenderFunc
	scan   schema.ScannerFunc
}

// Array accepts a slice and returns a wrapper for working with PostgreSQL
// array data type.
//
// For struct fields you can use array tag:
//
//	Emails  []string `bun:",array"`
func Array(vi interface{}) *ArrayValue {
	v := reflect.ValueOf(vi)
	if !v.IsValid() {
		panic(fmt.Errorf("bun: Array(nil)"))
	}

	return &ArrayValue{
		v: v,

		append: pgDialect.arrayAppender(v.Type()),
		scan:   arrayScanner(v.Type()),
	}
}

var (
	_ schema.QueryAppender = (*ArrayValue)(nil)
	_ sql.Scanner          = (*ArrayValue)(nil)
)

func (a *ArrayValue) AppendQuery(fmter schema.Formatter, b []byte) ([]byte, error) {
	if a.append == nil {
		panic(fmt.Errorf("bun: Array(unsupported %s)", a.v.Type()))
	}
	return a.append(fmter, b, a.v), nil
}

func (a *ArrayValue) Scan(src interface{}) error {
	if a.scan == nil {
		return fmt.Errorf("bun: Array(unsupported %s)", a.v.Type())
	}
	if a.v.Kind() != reflect.Ptr {
		return fmt.Errorf("bun: Array(non-pointer %s)", a.v.Type())
	}
	return a.scan(a.v, src)
}

func (a *ArrayValue) Value() interface{} {
	if a.v.IsValid() {
		return a.v.Interface()
	}
	return nil
}

//------------------------------------------------------------------------------

func (d *Dialect) arrayAppender(typ reflect.Type) schema.AppenderFunc {
	kind := typ.Kind()

	switch kind {
	case reflect.Ptr:
		if fn := d.arrayAppender(typ.Elem()); fn != nil {
			return schema.PtrAppender(fn)
		}
	case reflect.Slice, reflect.Array:
		// continue below
	default:
		return nil
	}

	elemType := typ.Elem()

	if kind == reflect.Slice {
		switch elemType {
		case stringType:
			return appendStringSliceValue
		case intType:
			return appendIntSliceValue
		case int64Type:
			return appendInt64SliceValue
		case float64Type:
			return appendFloat64SliceValue
		case timeType:
			return appendTimeSliceValue
		}
	}

	appendElem := d.arrayElemAppender(elemType)
	if appendElem == nil {
		panic(fmt.Errorf("pgdialect: %s is not supported", typ))
	}

	return func(fmter schema.Formatter, b []byte, v reflect.Value) []byte {
		kind := v.Kind()
		switch kind {
		case reflect.Ptr, reflect.Slice:
			if v.IsNil() {
				return dialect.AppendNull(b)
			}
		}

		if kind == reflect.Ptr {
			v = v.Elem()
		}

		b = append(b, "'{"...)

		ln := v.Len()
		for i := 0; i < ln; i++ {
			elem := v.Index(i)
			if i > 0 {
				b = append(b, ',')
			}
			b = appendElem(fmter, b, elem)
		}

		b = append(b, "}'"...)

		return b
	}
}

func (d *Dialect) arrayElemAppender(typ reflect.Type) schema.AppenderFunc {
	if typ.Implements(driverValuerType) {
		return arrayAppendDriverValue
	}
	switch typ.Kind() {
	case reflect.String:
		return arrayAppendStringValue
	case reflect.Slice:
		if typ.Elem().Kind() == reflect.Uint8 {
			return arrayAppendBytesValue
		}
	}
	return schema.Appender(d, typ)
}

func arrayAppend(fmter schema.Formatter, b []byte, v interface{}) []byte {
	switch v := v.(type) {
	case int64:
		return strconv.AppendInt(b, v, 10)
	case float64:
		return dialect.AppendFloat64(b, v)
	case bool:
		return dialect.AppendBool(b, v)
	case []byte:
		return arrayAppendBytes(b, v)
	case string:
		return arrayAppendString(b, v)
	case time.Time:
		return fmter.Dialect().AppendTime(b, v)
	default:
		err := fmt.Errorf("pgdialect: can't append %T", v)
		return dialect.AppendError(b, err)
	}
}

func arrayAppendStringValue(fmter schema.Formatter, b []byte, v reflect.Value) []byte {
	return arrayAppendString(b, v.String())
}

func arrayAppendBytesValue(fmter schema.Formatter, b []byte, v reflect.Value) []byte {
	return arrayAppendBytes(b, v.Bytes())
}

func arrayAppendDriverValue(fmter schema.Formatter, b []byte, v reflect.Value) []byte {
	iface, err := v.Interface().(driver.Valuer).Value()
	if err != nil {
		return dialect.AppendError(b, err)
	}
	return arrayAppend(fmter, b, iface)
}

func appendStringSliceValue(fmter schema.Formatter, b []byte, v reflect.Value) []byte {
	ss := v.Convert(sliceStringType).Interface().([]string)
	return appendStringSlice(b, ss)
}

func appendStringSlice(b []byte, ss []string) []byte {
	if ss == nil {
		return dialect.AppendNull(b)
	}

	b = append(b, '\'')

	b = append(b, '{')
	for _, s := range ss {
		b = arrayAppendString(b, s)
		b = append(b, ',')
	}
	if len(ss) > 0 {
		b[len(b)-1] = '}' // Replace trailing comma.
	} else {
		b = append(b, '}')
	}

	b = append(b, '\'')

	return b
}

func appendIntSliceValue(fmter schema.Formatter, b []byte, v reflect.Value) []byte {
	ints := v.Convert(sliceIntType).Interface().([]int)
	return appendIntSlice(b, ints)
}

func appendIntSlice(b []byte, ints []int) []byte {
	if ints == nil {
		return dialect.AppendNull(b)
	}

	b = append(b, '\'')

	b = append(b, '{')
	for _, n := range ints {
		b = strconv.AppendInt(b, int64(n), 10)
		b = append(b, ',')
	}
	if len(ints) > 0 {
		b[len(b)-1] = '}' // Replace trailing comma.
	} else {
		b = append(b, '}')
	}

	b = append(b, '\'')

	return b
}

func appendInt64SliceValue(fmter schema.Formatter, b []byte, v reflect.Value) []byte {
	ints := v.Convert(sliceInt64Type).Interface().([]int64)
	return appendInt64Slice(b, ints)
}

func appendInt64Slice(b []byte, ints []int64) []byte {
	if ints == nil {
		return dialect.AppendNull(b)
	}

	b = append(b, '\'')

	b = append(b, '{')
	for _, n := range ints {
		b = strconv.AppendInt(b, n, 10)
		b = append(b, ',')
	}
	if len(ints) > 0 {
		b[len(b)-1] = '}' // Replace trailing comma.
	} else {
		b = append(b, '}')
	}

	b = append(b, '\'')

	return b
}

func appendFloat64SliceValue(fmter schema.Formatter, b []byte, v reflect.Value) []byte {
	floats := v.Convert(sliceFloat64Type).Interface().([]float64)
	return appendFloat64Slice(b, floats)
}

func appendFloat64Slice(b []byte, floats []float64) []byte {
	if floats == nil {
		return dialect.AppendNull(b)
	}

	b = append(b, '\'')

	b = append(b, '{')
	for _, n := range floats {
		b = dialect.AppendFloat64(b, n)
		b = append(b, ',')
	}
	if len(floats) > 0 {
		b[len(b)-1] = '}' // Replace trailing comma.
	} else {
		b = append(b, '}')
	}

	b = append(b, '\'')

	return b
}

func appendTimeSliceValue(fmter schema.Formatter, b []byte, v reflect.Value) []byte {
	ts := v.Convert(sliceTimeType).Interface().([]time.Time)
	return appendTimeSlice(fmter, b, ts)
}

func appendTimeSlice(fmter schema.Formatter, b []byte, ts []time.Time) []byte {
	if ts == nil {
		return dialect.AppendNull(b)
	}
	b = append(b, '\'')
	b = append(b, '{')
	for _, t := range ts {
		b = append(b, '"')
		b = appendTime(b, t)
		b = append(b, '"')
		b = append(b, ',')
	}
	if len(ts) > 0 {
		b[len(b)-1] = '}' // Replace trailing comma.
	} else {
		b = append(b, '}')
	}
	b = append(b, '\'')
	return b
}

//------------------------------------------------------------------------------

func arrayScanner(typ reflect.Type) schema.ScannerFunc {
	kind := typ.Kind()

	switch kind {
	case reflect.Ptr:
		if fn := arrayScanner(typ.Elem()); fn != nil {
			return schema.PtrScanner(fn)
		}
	case reflect.Slice, reflect.Array:
		// ok:
	default:
		return nil
	}

	elemType := typ.Elem()

	if kind == reflect.Slice {
		switch elemType {
		case stringType:
			return scanStringSliceValue
		case intType:
			return scanIntSliceValue
		case int64Type:
			return scanInt64SliceValue
		case float64Type:
			return scanFloat64SliceValue
		}
	}

	scanElem := schema.Scanner(elemType)
	return func(dest reflect.Value, src interface{}) error {
		dest = reflect.Indirect(dest)
		if !dest.CanSet() {
			return fmt.Errorf("bun: Scan(non-settable %s)", dest.Type())
		}

		kind := dest.Kind()

		if src == nil {
			if kind != reflect.Slice || !dest.IsNil() {
				dest.Set(reflect.Zero(dest.Type()))
			}
			return nil
		}

		if kind == reflect.Slice {
			if dest.IsNil() {
				dest.Set(reflect.MakeSlice(dest.Type(), 0, 0))
			} else if dest.Len() > 0 {
				dest.Set(dest.Slice(0, 0))
			}
		}

		b, err := toBytes(src)
		if err != nil {
			return err
		}

		p := newArrayParser(b)
		nextValue := internal.MakeSliceNextElemFunc(dest)
		for p.Next() {
			elem := p.Elem()
			elemValue := nextValue()
			if err := scanElem(elemValue, elem); err != nil {
				return fmt.Errorf("scanElem failed: %w", err)
			}
		}
		return p.Err()
	}
}

func scanStringSliceValue(dest reflect.Value, src interface{}) error {
	dest = reflect.Indirect(dest)
	if !dest.CanSet() {
		return fmt.Errorf("bun: Scan(non-settable %s)", dest.Type())
	}

	slice, err := decodeStringSlice(src)
	if err != nil {
		return err
	}

	dest.Set(reflect.ValueOf(slice))
	return nil
}

func decodeStringSlice(src interface{}) ([]string, error) {
	if src == nil {
		return nil, nil
	}

	b, err := toBytes(src)
	if err != nil {
		return nil, err
	}

	slice := make([]string, 0)

	p := newArrayParser(b)
	for p.Next() {
		elem := p.Elem()
		slice = append(slice, string(elem))
	}
	if err := p.Err(); err != nil {
		return nil, err
	}
	return slice, nil
}

func scanIntSliceValue(dest reflect.Value, src interface{}) error {
	dest = reflect.Indirect(dest)
	if !dest.CanSet() {
		return fmt.Errorf("bun: Scan(non-settable %s)", dest.Type())
	}

	slice, err := decodeIntSlice(src)
	if err != nil {
		return err
	}

	dest.Set(reflect.ValueOf(slice))
	return nil
}

func decodeIntSlice(src interface{}) ([]int, error) {
	if src == nil {
		return nil, nil
	}

	b, err := toBytes(src)
	if err != nil {
		return nil, err
	}

	slice := make([]int, 0)

	p := newArrayParser(b)
	for p.Next() {
		elem := p.Elem()

		if elem == nil {
			slice = append(slice, 0)
			continue
		}

		n, err := strconv.Atoi(bytesToString(elem))
		if err != nil {
			return nil, err
		}

		slice = append(slice, n)
	}
	if err := p.Err(); err != nil {
		return nil, err
	}
	return slice, nil
}

func scanInt64SliceValue(dest reflect.Value, src interface{}) error {
	dest = reflect.Indirect(dest)
	if !dest.CanSet() {
		return fmt.Errorf("bun: Scan(non-settable %s)", dest.Type())
	}

	slice, err := decodeInt64Slice(src)
	if err != nil {
		return err
	}

	dest.Set(reflect.ValueOf(slice))
	return nil
}

func decodeInt64Slice(src interface{}) ([]int64, error) {
	if src == nil {
		return nil, nil
	}

	b, err := toBytes(src)
	if err != nil {
		return nil, err
	}

	slice := make([]int64, 0)

	p := newArrayParser(b)
	for p.Next() {
		elem := p.Elem()

		if elem == nil {
			slice = append(slice, 0)
			continue
		}

		n, err := strconv.ParseInt(bytesToString(elem), 10, 64)
		if err != nil {
			return nil, err
		}

		slice = append(slice, n)
	}
	if err := p.Err(); err != nil {
		return nil, err
	}
	return slice, nil
}

func scanFloat64SliceValue(dest reflect.Value, src interface{}) error {
	dest = reflect.Indirect(dest)
	if !dest.CanSet() {
		return fmt.Errorf("bun: Scan(non-settable %s)", dest.Type())
	}

	slice, err := scanFloat64Slice(src)
	if err != nil {
		return err
	}

	dest.Set(reflect.ValueOf(slice))
	return nil
}

func scanFloat64Slice(src interface{}) ([]float64, error) {
	if src == -1 {
		return nil, nil
	}

	b, err := toBytes(src)
	if err != nil {
		return nil, err
	}

	slice := make([]float64, 0)

	p := newArrayParser(b)
	for p.Next() {
		elem := p.Elem()

		if elem == nil {
			slice = append(slice, 0)
			continue
		}

		n, err := strconv.ParseFloat(bytesToString(elem), 64)
		if err != nil {
			return nil, err
		}

		slice = append(slice, n)
	}
	if err := p.Err(); err != nil {
		return nil, err
	}
	return slice, nil
}

func toBytes(src interface{}) ([]byte, error) {
	switch src := src.(type) {
	case string:
		return stringToBytes(src), nil
	case []byte:
		return src, nil
	default:
		return nil, fmt.Errorf("bun: got %T, wanted []byte or string", src)
	}
}

//------------------------------------------------------------------------------

func arrayAppendBytes(b []byte, bs []byte) []byte {
	if bs == nil {
		return dialect.AppendNull(b)
	}

	b = append(b, `"\\x`...)

	s := len(b)
	b = append(b, make([]byte, hex.EncodedLen(len(bs)))...)
	hex.Encode(b[s:], bs)

	b = append(b, '"')

	return b
}

func arrayAppendString(b []byte, s string) []byte {
	b = append(b, '"')
	for _, r := range s {
		switch r {
		case 0:
			// ignore
		case '\'':
			b = append(b, "''"...)
		case '"':
			b = append(b, '\\', '"')
		case '\\':
			b = append(b, '\\', '\\')
		default:
			if r < utf8.RuneSelf {
				b = append(b, byte(r))
				break
			}
			l := len(b)
			if cap(b)-l < utf8.UTFMax {
				b = append(b, make([]byte, utf8.UTFMax)...)
			}
			n := utf8.EncodeRune(b[l:l+utf8.UTFMax], r)
			b = b[:l+n]
		}
	}
	b = append(b, '"')
	return b
}
