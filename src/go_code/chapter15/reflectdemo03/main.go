package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v float64 = 1.2
	rVal := reflect.ValueOf(v)
	fmt.Printf("kind = %v Type = %v \n", rVal.Kind(), rVal.Type())

	iv := rVal.Interface()
	fmt.Printf("iv = %v iv = %T\n", iv, iv)

	num := iv.(float64)
	fmt.Printf("num = %v\n", num)
}