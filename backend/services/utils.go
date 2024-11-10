package services

import (
	"fmt"
	"reflect"
	"strings"
)

func GetStructMetadata(s interface{}) string {
	return getStructMetadata(reflect.TypeOf(s), "", 5)
}

func getStructMetadata(t reflect.Type, indent string, depth int) string {
	if depth < 0 {
		return ""
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem() // Get the element type if it's a pointer
	}

	result := fmt.Sprintf("%sStruct: %s\n", indent, t.Name())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldType := field.Type

		// Check if the field is a nested struct (excluding time.Time and other basic types)
		if fieldType.Kind() == reflect.Struct && !strings.Contains(fieldType.String(), "Time") {
			// Recursively call getStructMetadata for nested structs with increased indentation
			result += fmt.Sprintf("%sField: %s, Type:\n%s", indent, field.Name, getStructMetadata(fieldType, indent+"  ", depth-1))
		} else {
			// Use getFieldType for base types and other types (e.g., slices, maps, pointers)
			result += fmt.Sprintf("%sField: %s, Type: %s\n", indent, field.Name, getFieldType(fieldType, indent+"  ", depth-1))
		}
	}
	return result
}

func getFieldType(t reflect.Type, indent string, depth int) string {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "int"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "uint"
	case reflect.Float32, reflect.Float64:
		return "float"
	case reflect.Bool:
		return "bool"
	case reflect.String:
		return "string"
	case reflect.Slice:
		return "[]" + getFieldType(t.Elem(), indent+"  ", depth-1)
	case reflect.Array:
		return fmt.Sprintf("[%d]%s", t.Len(), getFieldType(t.Elem(), indent+"  ", depth-1))
	case reflect.Map:
		return fmt.Sprintf("map[%s]%s", getFieldType(t.Key(), indent+"  ", depth-1), getFieldType(t.Elem(), indent+"  ", depth-1))
	case reflect.Ptr:
		return "*" + getFieldType(t.Elem(), indent+"  ", depth-1)
	case reflect.Struct:
		return getStructMetadata(t, indent+"  ", depth-1)
	default:
		return t.Kind().String() // Fallback for unsupported types
	}
}
