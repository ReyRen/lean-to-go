package main

import (
	"fmt"
	"lean-to-go/factoryModel/model"
)

func main() {
	/*var stu = model.Student{
		Name:  "Tom",
		Score: 1,
	}

	fmt.Println(stu)*/

	var stu = model.NewStudent("Tom", 88.8)

	fmt.Println(*stu) //&{Tom 88.8}

	var stu2 = model.NewStudent2("Bob", 1.1)
	fmt.Println(stu2.GetScore())
}
