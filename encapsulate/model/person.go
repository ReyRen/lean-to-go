package model

import "fmt"

type person struct { // p cannot be used in other packages
	Name string
	age  int // cannot be used in other packages
	sal  float64
}

// factory
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// age
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("age input is out of range...")
	}
}
func (p *person) GetAge() int {
	return p.age
}

// sal
func (p *person) SetSal(sal float64) {
	if sal > 3000 && sal < 30000 {
		p.sal = sal
	} else {
		fmt.Println("sal input is out of range...")
	}
}
func (p *person) GetSal() float64 {
	return p.sal
}
