package main

import (
	"fmt"
	"time"
)

/*func main() {

	fmt.Println(time.Now())
	// create timer
	myTimer := time.NewTimer(time.Second * 2)

	nowTime := <- myTimer.C

	fmt.Println(nowTime)
}*/

/*func main() {
	time.Sleep(time.Second * 2)
}*/
/*
func main() {
	nowTime := <- time.After(time.Second)
	fmt.Println(nowTime)
}*/

func main() {
	myTimer := time.NewTimer(time.Second * 3)

	go func() {
		<-myTimer.C
		fmt.Println("goroutine timer done!")
	}()
	myTimer.Stop() // timer stop
	// 不会造成死锁，因为C的写端一直有操作系统把持着
	for {

	}
}
