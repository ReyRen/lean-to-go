package main

import "fmt"

//在defer将语句放入栈中时，也会将相关的值拷贝入栈

func sum(n1 int, n2 int) int {
	defer fmt.Println("defer n1=", n1)
	defer fmt.Println("defer n2=", n2)

	n1++
	n2++

	res := n1 + n2

	fmt.Println("sum res=", res)
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("main res=", res)

}

/*
sum res=32
defer n2=20
defer n1=10
main res=32
*/
