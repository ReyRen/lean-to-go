package main

import "fmt"

func calc(base int) (func(int) int, func(int) int) {
	fmt.Printf("%p\n", &base)

	add2 := func(i int) int {
		fmt.Printf("%p\n", &base)
		base += i
		return base
	}

	sub := func(i int) int {
		fmt.Printf("%p\n", &base)
		base -= i
		return base
	}

	return add2, sub
}

/*func main() {
	f1, f2 := calc(100)

	fmt.Println(f1(1), f2(2)) //执行顺序：f1 f2 println
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))
	fmt.Println(f1(7), f2(8))
}*/
