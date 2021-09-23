package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	//rVal.SetInt(20)
	rVal.Elem().SetInt(20) // like get the pointer pointed value
}

func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println(num)
}
