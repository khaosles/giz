package structs

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

/*
   @File: struct.go
   @Author: khaosles
   @Time: 2023/8/10 00:04
   @Desc:
*/

// defaultTagName is the default tag for struct fields to lookup.
var defaultTagName = "json"

// ToMap convert struct to map, only convert exported struct field
// map key is specified same as struct field tag `json` value.
func ToMap(v any) map[string]string {
	objValue := reflect.ValueOf(v)
	objType := objValue.Type()

	// 如果传入的不是结构体指针，则直接返回空 map
	if objType.Kind() == reflect.Ptr && objType.Elem().Kind() != reflect.Struct || objType.Kind() != reflect.Struct {
		return map[string]string{}
	}

	data := make(map[string]string)
	for i := 0; i < objValue.Elem().NumField(); i++ {
		field := objType.Elem().Field(i)
		value := objValue.Elem().Field(i)

		// 如果字段是空值，则跳过
		if reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface()) {
			continue
		}
		jsonTag := field.Tag.Get(defaultTagName)
		if jsonTag != "" {
			jsonTags := strings.Split(jsonTag, ",")
			// 忽略的json字段
			if len(jsonTags) > 1 && strings.TrimSpace(jsonTags[1]) == "-" {
				continue
			}
			name := strings.TrimSpace(jsonTags[0])
			data[name] = toString(value.Interface())
		} else {
			data[field.Name] = toString(value.Interface())
		}
		// 如果字段是结构体类型，则递归调用 StructToMap 进行转换
		if field.Type.Kind() == reflect.Struct {
			nestedData := ToMap(value.Interface())
			for k, v := range nestedData {
				data[k] = v
			}
		}
	}
	return data
}

func toString(value any) string {
	switch v := value.(type) {
	case string:
		return v
	default:
		b, _ := json.Marshal(v)
		return string(b)
	}
}

// ToMapInterface convert struct to map, only convert exported struct field
// map key is specified same as struct field tag `json` value.
func ToMapInterface(v interface{}) map[string]any {
	objValue := reflect.ValueOf(v)
	objType := objValue.Type()

	// 如果传入的不是结构体指针，则直接返回空 map
	if objType.Kind() != reflect.Ptr || objType.Elem().Kind() != reflect.Struct {
		return map[string]any{}
	}

	data := make(map[string]any)
	for i := 0; i < objValue.Elem().NumField(); i++ {
		field := objType.Elem().Field(i)
		value := objValue.Elem().Field(i)

		// 如果字段是空值，则跳过
		if reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface()) {
			continue
		}
		jsonTag := field.Tag.Get(defaultTagName)
		if jsonTag != "" {
			jsonTags := strings.Split(jsonTag, ",")
			// 忽略的json字段
			if len(jsonTags) > 1 && strings.TrimSpace(jsonTags[1]) == "-" {
				continue
			}
			name := strings.TrimSpace(jsonTags[0])
			data[name] = value.Interface()
		} else {
			data[field.Name] = value.Interface()
		}
		// 如果字段是结构体类型，则递归调用 StructToMap 进行转换
		if field.Type.Kind() == reflect.Struct {
			nestedData := ToMap(value.Interface())
			for k, v := range nestedData {
				data[k] = v
			}
		}
	}
	return data
}

// CopyProperties copy the common fields from struct `src` to struct `dst`
func CopyProperties[T, U any](dst *U, src T) error {
	// get the struct types and values using reflect
	srcType := reflect.TypeOf(src)
	srcValue := reflect.ValueOf(src)
	dstType := reflect.TypeOf(*dst)
	dstValue := reflect.ValueOf(dst).Elem()
	if dstValue.Kind() != reflect.Struct {
		return errors.New("CopyProperties: parameter dst should be struct pointer")
	}
	if srcValue.Kind() == reflect.Ptr {
		srcType = srcType.Elem()
		srcValue = srcValue.Elem()
	}
	if srcValue.Kind() != reflect.Struct {
		return errors.New("CopyProperties: parameter src should be struct pointer")
	}

	// Iterate over struct t1
	for i := 0; i < srcType.NumField(); i++ {
		// get filed name and value
		fieldName := srcType.Field(i).Name
		fieldValue := srcValue.Field(i).Interface()
		// wether has the same field in t2
		if _, ok := dstType.FieldByName(fieldName); ok {
			// set t2
			dstValue.FieldByName(fieldName).Set(reflect.ValueOf(fieldValue))
		}
	}
	return nil
}
