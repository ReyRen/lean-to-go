package main

import (
	"fmt"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("producer produce: ", i*i)
		out <- i * i
	}
	close(out) // or get deadlock
}

func consumer(in <-chan int) {
	for num := range in {
		//time.Sleep(time.Second * 2)
		fmt.Println("consumer get: ", num)
	}
}

func main() {
	//ch := make(chan int)
	ch := make(chan int)

	go producer(ch)

	consumer(ch)
}
