package main

import "fmt"

func add(base int) func(int) int {
	fmt.Printf("%p\n", &base)

	f := func(i int) int {
		fmt.Printf("inner %p\n", &base)
		base += i
		return base
	}

	return f
}

/*func main() {
	t1 := add(10)
	fmt.Println(t1(1), t1(2))

	t2 := add(1000)
	fmt.Println(t2(1), t2(2))
}*/
