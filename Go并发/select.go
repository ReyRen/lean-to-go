package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- 1
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true
		runtime.Goexit()
	}()

	for {
		select {
		case num := <-ch:
			fmt.Println(num)
		case <-quit:
			//break  break只能break出select
			return
		}
	}
}

/*
close(ch)要记得，因为如果不close，子routine写入完然后结束了，父routine(用break)中检测到一直在从一个不会再有写入但是没有close的channel，造成deadlock

chan即使close了，也是能read到0和false的
*/
