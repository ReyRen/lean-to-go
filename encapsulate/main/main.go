package main

import (
	"fmt"
	"lean-to-go/encapsulate/model"
)

func main() {
	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)

	fmt.Printf("%s %d %v", p.Name, p.GetAge(), p.GetSal())
}
