package utils

import (
	"encoding/json"
	"reflect"
)

func Struct2Json(obj interface{}) string {
	d, _ := json.Marshal(obj)
	return string(d)
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		j := t.Field(i).Tag.Get("json")
		if j != "-" {
			data[j] = v.Field(i).Interface()
		}
	}
	return data
}
