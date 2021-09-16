package main

import "fmt"

type Student struct {
}

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("No.%v paramater is bool type, the value is %v\n", index, x)
		case float32:
			fmt.Printf("No.%v paramater is float32 type, the value is %v\n", index, x)
		case int, int32, int64:
			fmt.Printf("No.%v paramater is int type, the value is %v\n", index, x)
		case string:
			fmt.Printf("No.%v paramater is string type, the value is %v\n", index, x)
		case Student:
			fmt.Printf("No.%v paramater is Student type, the value is %v\n", index, x)
		case *Student:
			fmt.Printf("No.%v paramater is *Student type, the value is %v\n", index, x)
		default:
			fmt.Printf("No.%v paramater type not sure, the value is %v\n", index, x)
		}
	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var name string = "tom"
	address := "Beijing"
	stu1 := Student{}
	stu2 := &Student{}

	TypeJudge(n1, n2, name, address, stu1, stu2)
}
