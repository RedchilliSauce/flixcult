package utils

import (
	"fmt"
	"net/url"
	"reflect"

	"github.com/labstack/echo"
)

const maxFilterSize int = 20

//FieldMatch ...
func FieldMatch(inpURL string, fieldname string, value []interface{}, p interface{}, t reflect.Type, l echo.Logger) (bool, error) {

	inpFilter, err := url.ParseQuery(inpURL)
	if err != nil {
		return false, err
	}
	if len(inpFilter) > maxFilterSize {
		//TODO error
		return false, err
	}

	v := reflect.ValueOf(p).Elem().Convert(t)
	typeOfV := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if fieldname == typeOfV.Field(i).Tag.Get("filter") {
			fv := fmt.Sprint(v.Field(i).Interface())
			inpv := fmt.Sprint(value)
			l.Debugf("Fieldmatcher: Service data %s  vs Input %s", fv, inpv)
			if inpv == fv {
				return false, nil
			}
		}
	}
	return false, nil
}
