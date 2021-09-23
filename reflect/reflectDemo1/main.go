package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	// get variable type by reflect
	//1. get reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp) // int

	//2. get reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)                        // 100 NOTE： this is not the int 100
	fmt.Printf("rVal=%v, rVal Type=%T\n", rVal, rVal) // rVal=100, rVal Type=reflect.Value
	n2 := 2 + rVal.Int()
	fmt.Println(n2)

	//3. rVal covert back to interface{}
	iV := rVal.Interface()
	//4. assert put interface to real type
	num2 := iV.(int)
	fmt.Printf("num2=%v, num2 Type=%T\n", num2, num2) //num2=100, num2 Type=int
}

type Student struct {
	Name string
	Age  int
}

func reflectTest02(b interface{}) {
	// get variable type by reflect
	//1. get reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp) //  main.Student

	//2. get reflect.Value
	rVal := reflect.ValueOf(b)
	iV := rVal.Interface()
	fmt.Printf("iV=%v, iV Type=%T\n", iV, iV) //iV={Tom 20}, iV Type=main.Student 虽然取出这样的值是Student结构体类型，但是要取出结构体中的值还是需要类型断言，因为编译阶段interface{}类型的编译器是找不到有相应结构体字段的，只有运行阶段，结构体才会去转换为对应的类型. 虽然取出这样的值是Student结构体类型，但是要取出结构体中的值还是需要类型断言，因为编译阶段interface{}类型的编译器是找不到有相应结构体字段的，只有运行阶段，结构体才会去转换为对应的类型.
	switch iV.(type) {
	case Student:
		fmt.Println(iV.(Student).Name)
	}
}

func reflectTest03(b interface{}) {
	// get variable type by reflect
	//1. get reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp) //  main.Student

	//2. get reflect.Value
	rVal := reflect.ValueOf(b)
	//3. Get variable kind
	//(1)rVal.Kind() ==>
	//(2)rTyp.Kind() ==>
	fmt.Printf("rVal kind=%v, rTyp kind=%v", rVal.Kind(), rTyp.Kind()) // struct struct

	iV := rVal.Interface()
	fmt.Printf("iV=%v, iV Type=%T\n", iV, iV) //iV={Tom 20}, iV Type=main.Student 虽然取出这样的值是Student结构体类型，但是要取出结构体中的值还是需要类型断言，因为编译阶段interface{}类型的编译器是找不到有相应结构体字段的，只有运行阶段，结构体才会去转换为对应的类型. 虽然取出这样的值是Student结构体类型，但是要取出结构体中的值还是需要类型断言，因为编译阶段interface{}类型的编译器是找不到有相应结构体字段的，只有运行阶段，结构体才会去转换为对应的类型.
	switch iV.(type) {
	case Student:
		fmt.Println(iV.(Student).Name)
	}
}

func main() {
	var num int = 100
	reflectTest01(num)

	stu := Student{
		Name: "Tom",
		Age:  20,
	}
	reflectTest02(stu)
}

/*
	reflect.Value.Kind获取变量的类别，返回的是一个常量
	Type和Kind的区别：
		Type是类型，Kind是类别。Type和Kind可能相同，也可能不同
		var num int = 10 num的kind和type都是int
		var stu Student stu的Type是pkg1.Student,kind是struct
*/
