package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex // new mutex is 0, it means none locked

func printer(str string) {
	mutex.Lock()
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
	mutex.Unlock()
}

func person1() {
	printer("hello")
}

func person2() {
	printer("world")
}

func main() {
	go person1()
	go person2()
	for true {

	}
}
