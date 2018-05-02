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

	// demo1: Object -> reflect -> Object
	demoFunc1(&v)

	// demo2: Type -> Object
	demoFunc2(&t)

	// demo3: get struct field
	printStructField(&t, &v)

	// demo4: modify Object
	modify(&t, &v)

	// demo5: get struct method
	callSetMethod(&v, "SetID", []interface{}{"demo2"})
	callGetMethod(&v, "GetID")

	// demo6: get struct tag
	getTag(&t, "ID", "json")
	getTag(&t, "Kind", "json")
	getTag(&t, "APIVersion", "json")
	getTag(&t, "ResourceVersion", "json")
	getTag(&t, "Name", "json")
}

// demo1: Object -> reflect -> Object
func demoFunc1(v *reflect.Value) {
	if s, ok := (*v).Interface().(reflectDemo.DemoJSONBase); ok {
		fmt.Printf("The DemonJSONBase is %s\n", *s.ID)
		fmt.Printf("The DemonJSONBase is %s\n", *s.Kind)
		fmt.Printf("The DemonJSONBase is %s\n", *s.APIVersion)
		fmt.Printf("The DemonJSONBase is %d\n", *s.ResourceVersion)
		fmt.Printf("The DemonJSONBase is %s\n", s.Name)
		fmt.Printf("The DemonJSONBase is %s\n", s.Value)
	} else {
		fmt.Println("error")
	}
}

// demo2: Type -> Object
func demoFunc2(t *reflect.Type) {
	v := reflect.New(*t)
	fmt.Println(v.Type().String())
}

// demo3: traverse struct field
func printStructField(t *reflect.Type, v *reflect.Value) {
	fieldNum := (*t).NumField()
	for i := 0; i < fieldNum; i++ {
		fv := (*v).Field(i)
		ft := (*t).Field(i)
		switch fv.Kind() {
		case reflect.String:
			fmt.Printf("The %d th %s types %s valuing %s\n", i, ft.Name, "string", fv.String())
		case reflect.Int:
			fmt.Printf("The %d th %s types %s valuing %d\n", i, ft.Name, "int", fv.Int())
		case reflect.Ptr:
			p := fv.Pointer()
			fmt.Printf("The %d th %s types %s valuing %v\n", i, ft.Name, "pointer", p)
		}
		fmt.Printf("field: %s\n", (*t).Field(i).Name)
	}
	fmt.Println("")
}

// demo4: modify Object
func modify(t *reflect.Type, v *reflect.Value) {
	for i := 0; i < (*v).NumField(); i++ {
		fv := (*v).Field(i)
		ft := (*t).Field(i)
		if !fv.CanSet() {
			fmt.Printf("The %d th %s is unaccessible\n", i, ft.Name)
		}
		switch fv.Kind() {
		case reflect.String:
			fv.SetString("test")
			fmt.Printf("string %s\n", ft.Name)
		case reflect.Int:
			fv.SetInt(18)
			fmt.Printf("int %s\n", ft.Name)
		case reflect.Ptr:
			fmt.Printf("ptr %v\n", ft.Name)
			continue
		}
	}
}

// demo5: call struct method
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

// demo6: get struct tag
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
