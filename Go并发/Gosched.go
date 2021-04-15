package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("this is goroutine test")
			//time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		runtime.Gosched()
		fmt.Println("this is main test")
		//time.Sleep(100 * time.Millisecond)
	}
}
