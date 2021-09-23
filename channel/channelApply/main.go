package main

import "fmt"

func writeData(intchan chan int) {
	for i := 0; i <= 50; i++ {
		intchan <- i
	}
	close(intchan)
}
func readData(intchan chan int, exitchan chan bool) {
	for {
		v, ok := <-intchan // 如果上面的channel不close，就会发生死锁，因为从一个空的chan读取数据会直接死锁
		if !ok {
			break
		}
		fmt.Printf("readData=%v\n", v)
	}
	exitchan <- true
	close(exitchan)
}

func main() {
	// create two channel
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
