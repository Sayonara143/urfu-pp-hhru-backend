package pagination

import (
	"context"
	"encoding/json"
)

type contextKeyType string

const key contextKeyType = "pagination"

func Set(ctx context.Context, p *Pagination) context.Context {
	return context.WithValue(ctx, key, p)
}

func Get(ctx context.Context) *Pagination {
	return ctx.Value(key).(*Pagination)
}

type Pagination struct {
	page       int
	perPage    int
	itemsTotal int
	pagesTotal int
}

func New() *Pagination {
	return &Pagination{page: 1, perPage: 10}
}

func (p *Pagination) SetPage(page int) {
	p.page = page
}

func (p *Pagination) SetPerPage(perPage int) {
	if perPage > 0 && perPage <= 1000 {
		p.perPage = perPage
	}
}

func (p *Pagination) Limit() int {
	return p.perPage
}

func (p *Pagination) ItemsTotal() int {
	return p.itemsTotal
}

func (p *Pagination) Offset() int {
	return (p.page - 1) * p.perPage
}

func (p *Pagination) SetItemsTotal(itemsTotal int) {
	p.itemsTotal = itemsTotal

	d := itemsTotal / p.perPage
	if itemsTotal%p.perPage > 0 {
		d += 1
	}
	p.pagesTotal = d
}

func (p *Pagination) PaginationResponse() json.RawMessage {
	r := map[string]int{
		"page":     p.page,
		"per-page": p.perPage,
	}
	if p.pagesTotal > 0 {
		r["pages-total"] = p.pagesTotal
		r["items-total"] = p.itemsTotal

		n := p.page + 1
		if n > p.pagesTotal {
			n = p.pagesTotal
		}
		if n != p.page {
			r["next-page"] = n
		}

		pr := p.page - 1
		if pr < 1 {
			pr = 1
		}
		if pr != p.page {
			r["prev-page"] = pr
		}
	}

	bb, _ := json.Marshal(r)
	return bb
}
