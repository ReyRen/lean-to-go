package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond

func producer(out chan<- int, idx int) {
	for {
		cond.L.Lock()       // if lock success, continue; or blocked here
		for len(out) == 3 { // len is left not read data
			cond.Wait()
			// 这里必须使用for循环。避免的问题是虚假唤醒。
			// 所谓虚假唤醒就是p1 p2 p3 p4已经全部都wait了，c1取走一个元素然后唤醒，此时唤醒了p1, p1从wait函数往下执行写入新的数据
			// 再次唤醒别的goroutine，这时候p2抢上了，也从它自己协程的wait开始执行，如果if语句，那么直接就向下执行了。
		}
		num := rand.Intn(10000)
		out <- num
		fmt.Printf("%dth producer, produced %3d, commonArea left %d data\n", idx, num, len(out))
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Second)
	}
}

func consumer(in <-chan int, idx int) {
	for {
		cond.L.Lock()
		for len(in) == 0 { // len is left not read data
			cond.Wait()
		}
		num := <-in
		fmt.Printf("%dth consumer, consume %3d, commonArea left %d data\n", idx, num, len(in))
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	quit := make(chan bool)

	commonArea := make(chan int, 3)
	cond.L = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go producer(commonArea, i+1)
	}

	for i := 0; i < 3; i++ {
		go consumer(commonArea, i+1)
	}

	<-quit
}
