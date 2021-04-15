package main

import (
	"fmt"
	"time"
)

func sing() {
	for i := 0; i < 5; i++ {
		fmt.Println("----------------singing--------------")
		time.Sleep(100 * time.Millisecond)
	}
}

func dance() {
	for i := 0; i < 5; i++ {
		fmt.Println("---------------dancing---------------")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sing()
	go dance()
	for true {

	}
}
