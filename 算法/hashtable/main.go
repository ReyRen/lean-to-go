package main

import (
	"fmt"
	"os"
)

type Employ struct {
	Id   int
	Name string
	Next *Employ
}

// head also have data
type EmployLinklist struct {
	Head *Employ
}

func (this *EmployLinklist) InsertLink(emp *Employ) {
	cur := this.Head
	var pre *Employ = nil
	if cur == nil {
		this.Head = emp
		return
	}
	for {
		if cur != nil {
			if cur.Id > emp.Id {
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}
	pre.Next = emp
	emp.Next = cur
}
func (this *EmployLinklist) ShowLink(hashIndex int) {
	if this.Head == nil {
		fmt.Printf("[%d] is empty\n", hashIndex)
	}
	cur := this.Head
	for {
		if cur != nil {
			fmt.Printf("[%d]-[name: %s employId: %d]--->", hashIndex, cur.Name, cur.Id)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}
func (this *EmployLinklist) FindById(id int) *Employ {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

type HashTable struct {
	LinkArr [7]EmployLinklist
}

func (this *HashTable) InsertHash(emp *Employ) {
	linkIndex := this.HashFunc(emp.Id)
	this.LinkArr[linkIndex].InsertLink(emp)
}
func (this *HashTable) ShowAllHashTable() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}
func (this *HashTable) HashFunc(id int) int {
	return id % 7
}
func (this *HashTable) FindById(id int) *Employ {
	linkIndex := this.HashFunc(id)
	return this.LinkArr[linkIndex].FindById(id)
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashTable HashTable
	for {
		fmt.Println("================Employment System================")
		fmt.Println("input(add employee)")
		fmt.Println("show(show employee)")
		fmt.Println("find(find employee)")
		fmt.Println("exit(exit system)")
		fmt.Println("Please input:")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("Please input employee ID:")
			fmt.Scanln(&id)
			fmt.Println("Please input employee NAME:")
			fmt.Scanln(&name)
			emp := &Employ{
				Id:   id,
				Name: name,
			}
			hashTable.InsertHash(emp)
		case "show":
			hashTable.ShowAllHashTable()
		case "find":
			fmt.Println("Please input employee ID:")
			fmt.Scanln(&id)
			emp := hashTable.FindById(id)
			if emp == nil {
				fmt.Println("Employ not exist..")
			} else {
				//get
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Input Error...")
		}
	}
}
