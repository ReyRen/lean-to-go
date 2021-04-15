package main

import "fmt"

func main() {
	/*ch := make(chan int)
	var sendCh chan <- int = ch
	sendCh <- 789

	var recvCh <- chan int = ch
	num := <- recvCh*/
	ch := make(chan int)
	go func() {
		send(ch)
	}()
	recv(ch)
}

func send(out chan<- int) { //传递的就是引用，这样保证使用的是同一个channel
	out <- 789
	close(out)
}

func recv(in <-chan int) {
	n := <-in
	fmt.Println("read get: ", n)
}
