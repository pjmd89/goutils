package dbutils

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type Tags struct {
	Name        string
	Default     string
	IsDefault   bool
	IsID        bool
	IsObjectID  bool
	IsOmit      bool
	IsNested    bool
	CreatedDate bool
	UpdatedDate bool
	Change      string
	Compose     []string
}

func CreateStruct(instance interface{}, update bool) (r interface{}) {
	valueOf := reflect.ValueOf(instance)
	typeOf := valueOf.Type()
	structFields := make([]reflect.StructField, 0)
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)

		tag := fmt.Sprintf("%v", typeOf.Field(i).Tag)
		tagFind := regexp.MustCompile(`bson:"[^"\-]+"`)
		notFind := regexp.MustCompile(`omitempty`)
		result := tagFind.FindString(tag)

		if !notFind.MatchString(result) && strings.Trim(result, " ") != "" {
			replace := regexp.MustCompile(`(bson:"[^"]+)(["])`)
			tag = replace.ReplaceAllString(tag, "$1,omitempty\"")
		}
		switch field.Type.Kind() {
		case reflect.Struct:
			newInstance := reflect.New(field.Type).Elem().Interface()
			field.Type = reflect.TypeOf(CreateStruct(newInstance, update))
		case reflect.Ptr:
			if field.Type.Elem().Kind() == reflect.Struct {
				newPtrInstance := reflect.New(field.Type.Elem()).Elem().Interface()
				newInstance := CreateStruct(newPtrInstance, update)
				field.Type = reflect.New(reflect.TypeOf(newInstance)).Type()
			}
		case reflect.Slice, reflect.Array:

		}
		if update {
			replaceGQL := regexp.MustCompile(`(gql:"[^"]+)(["])`)
			result := replaceGQL.FindString(tag)

			if result != "" {
				replace := regexp.MustCompile(`:`)
				result2 := replace.Split(result, -1)
				tag2 := strings.Replace(result2[1], `"`, "", -1)
				replace2 := regexp.MustCompile(`,`)
				result3 := replace2.Split(tag2, -1)
				replace3 := regexp.MustCompile(`=`)
				var updateTag []string
				for _, sv := range result3 {
					result4 := replace3.Split(sv, -1)
					if result4[0] != "default" {
						updateTag = append(updateTag, sv)
					}
				}
				tag = replaceGQL.ReplaceAllString(tag, `gql:"`+strings.Join(updateTag, ",")+`"`)
			}
			field.Tag = reflect.StructTag(tag)
		}
		structFields = append(structFields, field)
	}
	newType := reflect.StructOf(structFields)
	newStruct := reflect.New(newType).Elem().Interface()
	if update {
		//newStruct = reflect.New(newType).Elem().Convert(typeOf).Interface()
	}
	r = newStruct
	return r
}
func GetTags(field reflect.StructField) (r Tags) {
	tag := field.Tag.Get("gql")
	if tag != "" {
		tagSplit := strings.Split(tag, ",")
		for _, v := range tagSplit {
			dataSplit := strings.Split(v, "=")
			switch dataSplit[0] {
			case "name":
				r.Name = dataSplit[1]
				break
			case "default":
				r.Default = dataSplit[1]
				r.IsDefault = true
				break
			case "id":
				isTrue := false
				if strings.Trim(dataSplit[1], " ") == "true" {
					isTrue = true
				}
				r.IsID = isTrue
				break
			case "objectID":
				isTrue := false
				if strings.Trim(dataSplit[1], " ") == "true" {
					isTrue = true
				}
				r.IsObjectID = isTrue
				break
			case "nested":
				isTrue := false
				if strings.Trim(dataSplit[1], " ") == "true" {
					isTrue = true
				}
				r.IsNested = isTrue
				break
			case "omit":
				isTrue := false
				if strings.Trim(dataSplit[1], " ") == "true" {
					isTrue = true
				}
				r.IsOmit = isTrue
				break
			case "created":
				isTrue := false
				if strings.Trim(dataSplit[1], " ") == "true" {
					isTrue = true
				}
				r.CreatedDate = isTrue
				break
			case "updated":
				isTrue := false
				if strings.Trim(dataSplit[1], " ") == "true" {
					isTrue = true
				}
				r.UpdatedDate = isTrue
				break
			case "change":
				r.Change = dataSplit[1]
			case "compose":
				r.Compose = strings.Split(dataSplit[1], "|")
			}
		}
	}
	return r
}
