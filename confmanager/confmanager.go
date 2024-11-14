package confmanager

import (
	"errors"
	"io/ioutil"
	"reflect"

	"gopkg.in/yaml.v3"
)

func FromYAML(config interface{}, pathToFile string) error {
	t := reflect.TypeOf(config)
	if t.Kind() != reflect.Ptr {
		return errors.New("config can only assign values with pointer to struct or map")
	}

	fileData, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		return errors.New("failed to read config file " + pathToFile + ": " + err.Error())
	}

	err = yaml.Unmarshal(fileData, config)
	if err != nil {
		return errors.New("failed to parse config file: " + err.Error())
	}

	return nil
}
