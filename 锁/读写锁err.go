package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutexerr sync.RWMutex

func main() {

	rand.Seed(time.Now().UnixNano())

	ch := make(chan int) // used to data transfer

	for i := 0; i < 5; i++ { // 创建了5个读go程
		go readGoerr(ch, i+1) // 一个go程拿上读锁，然后其他的写go程全部都阻塞在拿锁上，然后造成整个程序channel不是读写同时在线而死锁
	}

	for i := 0; i < 5; i++ {
		go writeGoerr(ch, i+1)
	}

	for true {

	}
}

func writeGoerr(out chan<- int, index int) {
	num := rand.Intn(1000)
	rwMutexerr.Lock()
	out <- num
	fmt.Printf("%dth WRITE : %d\n", index, num)
	time.Sleep(time.Millisecond * 300)
	rwMutexerr.Unlock()
}

func readGoerr(in <-chan int, index int) {
	rwMutexerr.RLock()
	num := <-in
	fmt.Printf("%dth READ : %d\n", index, num)
	rwMutexerr.RUnlock()
}
