package utils

import "reflect"

// @description: 利用反射将结构体转化为map
// @example: map[Name:John Age:30] // 传入的结构体为：type Person struct { Name string `mapstructure:"Name"` Age int `mapstructure:"Age"` }
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ { // 遍历结构体的所有字段
		if obj1.Field(i).Tag.Get("mapstructure") != "" { // 如果字段有mapstructure标签，则使用标签的值作为key
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface() // 将字段的值作为value
		} else { // 字段名不变
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

// @description: 接受任何类型的参数 in，并返回一个指向相同类型的指针
func Pointer[T any](in T) (out *T) {
	return &in
}
