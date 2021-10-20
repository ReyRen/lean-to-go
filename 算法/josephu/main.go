package main

import (
	"fmt"
)

//
//  Boy
//  @Description: node struct
//
type Boy struct {
	Id   int
	Next *Boy
}

//
//  GenerageSingleCircleLinkList
//  @Description: make single circle linklist
//  @param boyNum: number of node in circle linklist
//  @return *Boy: the first node pointer in circle linklist
//
func GenerageSingleCircleLinkList(boyNum int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}

	if boyNum < 1 {
		fmt.Println("boyNum too small...")
		return first
	}

	for i := 1; i <= boyNum; i++ {
		boy := &Boy{
			Id: i,
		}
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}
	return first
}

//
//  ShowBoy
//  @Description: show all single circle linklist node
//  @param first
//
func ShowBoy(first *Boy) {
	if first.Next == nil {
		fmt.Println("linklist is empty...")
		return
	}

	curBoy := first
	for {
		fmt.Printf("Boy(%d)--->", curBoy.Id)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
	fmt.Printf("Boy(%d)", first.Id) // rear
}

//
//  PlayGame
//  @Description:
//  @param first: head
//  @param startId: k
//  @param countNum: step
//
func PlayGame(first *Boy, startId int, countNum int) {
	if first.Next == nil {
		fmt.Println("empty linklist")
		return
	}

	// tail: used to delete node in linklist
	tail := first
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}
	// move first to (startId - 1) why -1? first already in 1
	for i := 1; i <= startId-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	fmt.Println()
	for {
		ShowBoy(first)
		fmt.Println()
		// count to countNum - 1
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("Boy(%d) out\n", first.Id)
		first = first.Next
		tail.Next = first
		if tail == first {
			break
		}
	}
	fmt.Printf("Boy(%d) out\n", first.Id)
}

func main() {
	first := GenerageSingleCircleLinkList(100)
	//ShowBoy(first)
	PlayGame(first, 20, 31)
}
