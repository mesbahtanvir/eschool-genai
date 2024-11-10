package services

import (
	"fmt"
	"reflect"
)

func GetStructMetadata(s interface{}) string {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem() // Get the element type if it's a pointer
	}

	result := fmt.Sprintf("Struct: %s\n", t.Name())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		result += fmt.Sprintf("Field: %s, Type: %s\n", field.Name, field.Type)
	}
	return result
}
