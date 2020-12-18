package main

import "fmt"

func main() {
	//判断map中的key是否存在

	var m map[int]string = map[int]string{1: "111", 2: "222"}
	if value, exist := m[1]; exist {
		//m[1] 返回两个值，第一个是value, 第二个是bool代表key是否存在
		fmt.Println(value)
	}
}
