package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println(num)
			case <-time.After(time.Second * 3):
				quit <- true
				return
			}
		}
	}()

	for i := 0; i < 2; i++ {
		ch <- i
		time.Sleep(time.Second * 2)
	}

	<-quit
	fmt.Println("finished!")
}
