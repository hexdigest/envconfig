package env_config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type InvalidSpecificationError struct{}

func (e InvalidSpecificationError) Error() string {
	return "env_config: invalid specification type must be a struct"
}

func Process(prefix string, spec interface{}) error {
	s := reflect.ValueOf(spec).Elem()
	if s.Kind() != reflect.Struct {
		return &InvalidSpecificationError{}
	}
	typeOfSpec := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if f.CanSet() {
			key := fmt.Sprintf("%s_%s", prefix, typeOfSpec.Field(i).Name)
			value := os.Getenv(strings.ToUpper(key))
			switch f.Kind() {
			case reflect.String:
				f.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				intValue, err := strconv.ParseInt(value, 0, f.Type().Bits())
				if err != nil {
					return nil
				}
				f.SetInt(intValue)
			}
		}
	}
	return nil
}
