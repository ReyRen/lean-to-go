package main

import "fmt"

type Person3 struct {
	age   int
	name  string
	slice []string
}

func initFunc() *Person3 {
	p := new(Person3)

	p.age = 1
	p.name = "x"
	p.slice = append(p.slice, "fucker")
	return p // 不能返回局部变量的地址值，可以返回局部变量的值
}

func main() {
	p2 := initFunc()
	fmt.Println(p2)
}
