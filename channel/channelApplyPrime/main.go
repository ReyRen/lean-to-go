package main

import (
	"fmt"
	"time"
)

func putNum(intchan chan int) {
	for i := 0; i < 800000; i++ {
		intchan <- i
	}
	close(intchan)
}

func primeNum(intchan chan int, primechan chan int, exitchan chan bool) {

	var flag bool
	for {
		nums, ok := <-intchan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < nums; i++ {
			if nums%i == 0 {
				//not prime
				flag = false
				break
			}
		}
		if flag {
			primechan <- nums
		}
	}
	fmt.Println("one of primeNum goroutine cannot get data, exit!")
	exitchan <- true
}

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)
	start := time.Now().UnixNano()
	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		end := time.Now().UnixNano()
		fmt.Printf("consume time using gorouting...%v\n", end-start)
		close(primeChan)
	}()

	for {
		//res, ok := <- primeChan
		_, ok := <-primeChan
		if !ok {
			break
		}
		//fmt.Printf("prime=%d\n", res)
	}

	close(exitChan)
}
