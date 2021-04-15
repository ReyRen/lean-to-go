package main

import (
	"fmt"
	"time"
)

//global chan
var channel = make(chan int)

func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(300 * time.Millisecond)
	}
}

func person1() {
	printer("hello")
	channel <- 111
}

func person2() {
	<-channel
	printer("world")
}

func main() {
	go person1()
	go person2()
	for {

	}
}
