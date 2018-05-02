package main

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/yangyumo123/k8s-demo/pkg/reflectDemo"
)

func main() {
	fmt.Println("")

	// create DemoJSONBase object
	json := reflectDemo.DemoJSONBase{}
	id := "demo"
	kind := "Container"
	apiVersion := "v1"
	resourceVersion := uint64(1)
	json.ID = &id
	json.Kind = &kind
	json.APIVersion = &apiVersion
	json.ResourceVersion = &resourceVersion
	json.Name = "username"
	json.Value = "123456"

	// struct type info
	t := reflect.TypeOf(json)

	// value info
	v := reflect.ValueOf(json)

	// demo1: get struct field
	printStructField(&t)

	// demo2: get struct method
	callSetMethod(&v, "SetID", []interface{}{"demo2"})
	callGetMethod(&v, "GetID")

	// demo3: get struct tag
	getTag(&t, "ID", "json")
	getTag(&t, "Kind", "json")
	getTag(&t, "APIVersion", "json")
	getTag(&t, "ResourceVersion", "json")
	getTag(&t, "Name", "json")
}

// demo1: get struct field name
func printStructField(t *reflect.Type) {
	fieldNum := (*t).NumField()
	for i := 0; i < fieldNum; i++ {
		fmt.Printf("field: %s\n", (*t).Field(i).Name)
	}
	fmt.Println("")
}

// demo2: call struct method
func callSetMethod(v *reflect.Value, method string, params []interface{}) {
	f := (*v).MethodByName(method)
	if f.IsValid() {
		args := make([]reflect.Value, len(params))
		for k, param := range params {
			args[k] = reflect.ValueOf(param)
		}
		// call
		f.Call(args)
	} else {
		fmt.Println("can't call " + method)
	}
	fmt.Println("")
}

func callGetMethod(v *reflect.Value, method string) {
	f := (*v).MethodByName(method)
	if f.IsValid() {
		args := []reflect.Value{}
		ret := f.Call(args)
		if ret[0].Kind() == reflect.String {
			fmt.Printf("GetID %s\n", ret[0].String())
		}
	} else {
		fmt.Println("can't call " + method)
	}
	fmt.Println("")
}

// demo3: get struct tag
func getTag(t *reflect.Type, field string, tagName string) {
	var (
		tagVal string
		err    error
	)
	fieldVal, ok := (*t).FieldByName(field)
	if ok {
		tagVal = fieldVal.Tag.Get(tagName)
	} else {
		err = errors.New("no field named:" + field)
	}
	fmt.Printf("get struct[%s] tag[%s]: %s, error:%v\n", field, tagName, tagVal, err)
	fmt.Println("")
}
