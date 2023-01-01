package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A string
}

func (t T) String() string {
	return t.A + "kate"
}

func main() {
	fmt.Println("--------------Map的反射--------------")
	/*
		the type of m is : map[string]string
		the value of m is : map[k1:v1 k2:v2 k3:v3]
	*/
	m := make(map[string]string, 10)
	m["k1"] = "v1"
	m["k2"] = "v2"
	m["k3"] = "v3"
	t := reflect.TypeOf(m)
	fmt.Println("the type of m is :", t)
	v := reflect.ValueOf(m)
	fmt.Println("the value of m is :", v)
	fmt.Println("--------------结构体的反射--------------")
	/*
		Field 0:bbbb
		Method 0: 0x889be0
		[bbbbkate]
	*/
	myStruct := T{A: "bbbb"}
	tt := reflect.ValueOf(myStruct)
	for i := 0; i < tt.NumField(); i++ {
		fmt.Printf("Field %d:%v\n", i, tt.Field(i))
	}
	for i := 0; i < tt.NumMethod(); i++ {
		fmt.Printf("Method %d: %v\n", i, tt.Method(i))
	}
	result := tt.Method(0).Call(nil)
	fmt.Println(result)
}
