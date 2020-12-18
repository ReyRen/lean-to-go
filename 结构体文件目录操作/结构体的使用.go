package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func main() {
	var man Person = Person{
		name: "andy",
		sex:  'm',
		age:  1,
	}
	fmt.Println(man.name)

	man2 := Person{sex: 's', age: 1}
	fmt.Println(man2.name)

	var man3 Person
	man3.name = "ss"
	man3.age = 99
}
