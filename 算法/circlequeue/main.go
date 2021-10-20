package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int
	array   [4]int
	head    int
	tail    int
}

func (this *CircleQueue) PushQueue(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue is full")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}
func (this *CircleQueue) PopQueue() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}
func (this *CircleQueue) ListQueue() {
	fmt.Println("Circle Queue:")
	size := this.Size()
	if size == 0 {
		fmt.Println("queue is empty")
	}

	tmpHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tmpHead, this.array[tmpHead])
		tmpHead = (tmpHead + 1) % this.maxSize
	}
	fmt.Println()
}

func (this *CircleQueue) IsFull() bool {
	//return (this.tail + 1) % this.maxSize == this.head
	return (this.tail+1)%this.maxSize == this.head
}
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}
func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main() {

	queue := &CircleQueue{
		maxSize: 4,
		head:    0,
		tail:    0,
	}

	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.PushQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Add ok")
			}
		case "get":
			val, err := queue.PopQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Get: ", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
