package systemutils

import "reflect"

func Struct2Map(s interface{}) map[string]interface{} {
	return structToMapRecursive(s)
}

func structToMapRecursive(s interface{}) map[string]interface{} {
	sValue := reflect.ValueOf(s)
	sType := reflect.TypeOf(s)

	if sType.Kind() == reflect.Ptr {
		sValue = sValue.Elem()
		sType = sType.Elem()
	}

	r := make(map[string]interface{})

	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)
		fieldValue := sValue.Field(i)

		if field.Type.Kind() == reflect.Struct {
			r[field.Name] = structToMapRecursive(fieldValue.Interface())
		} else {
			r[field.Name] = fieldValue.Interface()
		}
	}

	return r
}
