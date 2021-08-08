package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{1, 3, 6}
	aI := interface{}(a)
	fmt.Printf("TypeOf:\t %v \n", reflect.TypeOf(aI))
	fmt.Printf("Kind compare:\t %v \n", reflect.TypeOf(aI).Kind() == reflect.Int)
	fmt.Printf("ValueOf:\t %v \n", reflect.ValueOf(aI))
	switch reflect.TypeOf(aI).Kind() {
	case reflect.Int:
		fmt.Printf("Kind:\t int \n")
	case reflect.Float64:
		fmt.Printf("Kind:\t float64 \n")
	case reflect.Slice:
		fmt.Printf("Kind:\t slice \n")
	case reflect.Map:
		fmt.Printf("Kind:\t map \n")
	default:
		fmt.Printf("Kind:\t unknown \n")
	}

	var i int
	v := reflect.ValueOf(&i).Elem()
	fmt.Printf("CanSet:\t %t \n", v.CanSet())
	v.Set(reflect.ValueOf(200))
	fmt.Printf("v.Set:\t %v \n", v)
	i2 := v.Interface()
	if i3, ok := i2.(int); ok {
		fmt.Printf("v.Interface:\t type=%T value=%[1]v \n", i2)
		fmt.Printf("v.Interface.(int):\t type=%T value=%[1]v \n", i3)
	}

	rt := reflect.ArrayOf(2, reflect.TypeOf(0))
	fmt.Printf("rt Type: %v Kind: %v \n", rt, rt.Kind())
	fmt.Printf("rt.Len:\t %v \n", rt.Len())
	rv := reflect.New(rt).Elem()
	fmt.Printf("rv.Index(0):\t %v\n", rv.Index(0))
	rv.Index(0).SetInt(100)
	fmt.Printf("rv.Index(0).SetInt(100) rv.Index(0):\t%v \n", rv.Index(0))
	ary := [1]int{0}
	reflect.ValueOf(&ary).Elem().Index(0).SetInt(500)
	fmt.Printf("reflect.ValueOf(&ary).Elem().Index(0).SetInt(500):\t%v \n", ary)

	type User struct {
		Name string
		Age  int
	}

	u := User{}
	fmt.Printf("reflect.Type(User{}):\t%v \n",reflect.TypeOf(u))
	fmt.Printf("reflect.TypeOf(User{}).Kind() == reflect.Struct:\t%v \n", reflect.TypeOf(User{}).Kind() == reflect.Struct)
	t := reflect.StructOf([]reflect.StructField{
		reflect.StructField{Name: "Name", Type: reflect.TypeOf("")},
		reflect.StructField{Name: "Age", Type: reflect.TypeOf(0)},
	})
	fmt.Printf("%v\n", t)
	fmt.Printf("%t\n", reflect.TypeOf(u).Kind() == t.Kind())
}
