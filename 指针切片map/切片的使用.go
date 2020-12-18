package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 数组
	s := arr[2:5]                                 // {3, 4, 5}
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s)) // 3
	fmt.Println("cap(s) = ", cap(s)) // 8

	s2 := s[2:7]
	fmt.Println("s2 = ", s2)           // [5 6 7 8 9]
	fmt.Println("len(s2) = ", len(s2)) // 5
	fmt.Println("cap(s2) = ", cap(s2)) // 6
}

/*
解释：

切片的结构体中的三个字段分别是指向底层数组的指针、切片访问的元素的个数（长度）切片允许增长到的元素个数（容量）
所以能想象arr[2:5]的时候，第一个元素数组指针发生了移动，所以其cap也会相应的变化
*/
