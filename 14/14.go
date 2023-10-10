package main

import (
	"fmt"
	"reflect"
)

func defineVarType(object any) string {
	value := reflect.ValueOf(object)
	var t string
	switch value.Kind() {
	case reflect.Int:
		t = "int"
	case reflect.String:
		t = "string"
	case reflect.Bool:
		t = "bool"
	case reflect.Chan:
		t = "chan"
	default:
		t = "unknown type"
	}
	return t
}

func main() {
	arr := []any{5, "qwe", true, make(chan int), 3.5}
	for _, v := range arr {
		fmt.Printf("%v %s\n", v, defineVarType(v))
	}
}
