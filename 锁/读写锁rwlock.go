package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex sync.RWMutex
var value int

func main() {

	rand.Seed(time.Now().UnixNano())

	//ch := make(chan int) // used to data transfer

	for i := 0; i < 5; i++ {
		go readGo(i + 1)
	}

	for i := 0; i < 5; i++ {
		go writeGo(i + 1)
	}

	for true {

	}
}

func writeGo(index int) {
	num := rand.Intn(1000)
	rwMutex.Lock()
	value = num
	fmt.Printf("%dth WRITE : %d\n", index, num)
	time.Sleep(time.Millisecond * 300)
	rwMutex.Unlock()
}

func readGo(index int) {
	rwMutex.RLock()
	num := value
	fmt.Printf("%dth READ : %d\n", index, num)
	rwMutex.RUnlock()
}
