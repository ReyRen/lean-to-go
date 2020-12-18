package main

import "fmt"

func main() {
	var a int = 10

	var p *int = &a

	a = 100

	*p = 250

	fmt.Println("a=", a)

	// heap
	p = new(int)
}
