package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// boolean example
	// var b bool
	// fmt.Println(b)
	// b = true
	// fmt.Println(b)

	// view datatypes
	// var s string = "string"
	// var i int = 10
	// var f float32 = 1.2
	// fmt.Println(reflect.TypeOf(s))
	// fmt.Println(reflect.TypeOf(i))
	// fmt.Println(reflect.TypeOf(f))

	// convert data types
	var b bool
	fmt.Println(reflect.TypeOf(b))
	var s string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))
}
