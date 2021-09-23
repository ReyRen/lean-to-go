package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {

	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "Tom Jack"
	cat := Cat{
		Name: "xiaohuamao",
		Age:  100,
	}
	allChan <- cat

	<-allChan
	<-allChan

	newCat := <-allChan

	fmt.Printf("newCat=%T, newCat=%v\n", newCat, newCat) // newCat=main.Cat, newCat={xiaohuamao 100}
	//fmt.Printf("newCat.Name=%v\n", newCat.Name) // error
	/*
		这样的原因是：
		在编译阶段newCat是interface{}类型，运行阶段newCat是main.Cat类型
	*/
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v\n", a.Name)
}
