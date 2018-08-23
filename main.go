package main

import (
	"reflect"
	"fmt"
	"unicode"
)

type Request struct {
	Id int	`json:"id"`
	Name string	`json:"name"`
	inner string
	Age *int 	`json:"age"`
	Xi interface{}
}

func MyJsonEncode(obj interface{}) {
	var (
		i int

		objType reflect.Type
		objValue reflect.Value

		field reflect.StructField
		fieldValue reflect.Value

		fieldName string
	)

	// 接口是空(没装任何东西的interface{})
	if obj == nil {
		fmt.Println("空接口")
		return
	}

	// 反射变量
	objType = reflect.TypeOf(obj) // 反射类型
	objValue = reflect.ValueOf(obj) // 反射值

	// 如果是指针, 需要取值
	if objType.Kind() == reflect.Ptr {
		if objValue.IsNil() {		// 空指针
		fmt.Println("空指针")
			return
		}
		objType = objType.Elem() // 相当于类型为*ptr
		objValue = objValue.Elem() // 相当于值为*ptr
	}

	// 如果不是结构体, 则不需要递归处理
	if objType.Kind() != reflect.Struct {
		fmt.Println("普通值", objValue.Interface())
		return
	}

	// 递归处理结构体中的字段
	for i = 0; i < objType.NumField(); i++ {
		field = objType.Field(i)	// 获取字段类型
		fieldValue = objValue.Field(i) // 获取字段的值

		// 小写字段不导出
		fieldName = field.Name
		if unicode.IsLower(rune(fieldName[0])) {
			continue
		}

		// 打印这个字段的信息
		fmt.Println("字段:", field.Name, "类型:", field.Type, "标签:", field.Tag)

		// 递归编码这个字段
		MyJsonEncode(fieldValue.Interface())
	}
}


func main() {
	var (
		req *Request
	)

	req = &Request{
		Id: 1,
		Name: "owen",
	}

	MyJsonEncode(req)
}
