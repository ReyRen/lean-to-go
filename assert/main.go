package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{
		x: 1,
		y: 2,
	}

	a = point

	var b Point
	b = a.(Point) // 如果想让一个空接口重新转换成既定类型，那么就需要使用类型断言
	fmt.Println(b)

	var x interface{}
	var b2 float64
	x = b2
	y, ok := x.(float64)
	if ok {
		fmt.Println("convert success!")
	} else {
		fmt.Println("convert err!")
	}
	fmt.Println("continue...", y)

}
