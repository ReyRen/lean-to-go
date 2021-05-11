package main

import (
	"fmt"
	"math/rand"
	"time"
)

//var valueChannel int

func main() {

	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go readGoChannel(ch, i+1)
	}

	for i := 0; i < 5; i++ {
		go writeGoChannel(ch, i+1)
	}

	for true {

	}
}

func writeGoChannel(out chan<- int, index int) {
	num := rand.Intn(1000)
	rwMutex.Lock()
	out <- num
	fmt.Printf("%dth WRITE : %d\n", index, num)
	time.Sleep(time.Millisecond * 300)
	rwMutex.Unlock()
}

func readGoChannel(in <-chan int, index int) {
	rwMutex.RLock()
	num := <-in
	fmt.Printf("%dth READ : %d\n", index, num)
	rwMutex.RUnlock()
}
