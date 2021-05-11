package main

import (
	"fmt"
	"time"
)

func main() {

	quit := make(chan bool)

	fmt.Println("now: ", time.Now())
	myTicker := time.NewTicker(time.Second)

	i := 0
	go func() {
		for {
			nowTime := <-myTicker.C
			i++
			fmt.Println("now time: ", nowTime)
			if i == 8 {
				quit <- true
				break
			}
		}
	}()
	<-quit
}
