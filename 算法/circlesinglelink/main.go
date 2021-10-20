package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head // circled with self
		fmt.Println(newCatNode, "joined into circle linklist")
		return
	}

	tmp := head
	for {
		if tmp.next == head {
			break
		}
		tmp = tmp.next
	}
	tmp.next = newCatNode
	newCatNode.next = head
}

func ListCatNode(head *CatNode) {

	tmp := head
	if tmp.next == nil {
		fmt.Println("empty list...")
		return
	}

	for {
		fmt.Printf("Cat Info : [id=%d name=%s] --->\n", tmp.no, tmp.name)
		if tmp.next == head {
			break
		}
		tmp = tmp.next
	}
}

func DelCatNode(head *CatNode, id int) *CatNode {
	tmp := head
	helper := head

	if tmp.next == nil {
		fmt.Println("empty linklist")
		return head
	}

	if tmp.next == head { // one head circle linklist
		if tmp.no == id {
			tmp.next = nil
		}
		return head
	}

	// not only have head
	for {
		if helper.next == head {
			break
		}
		helper = helper.next // point to the end node
	}

	flag := true
	for {
		if tmp.next == head {
			break
		}
		if tmp.no == id {
			if tmp == head { // delete head node
				head = head.next
			}
			helper.next = tmp.next
			fmt.Printf("Cat=%d\n", id)
			flag = false
			break
		}
		tmp = tmp.next
		helper = helper.next
	}

	if flag {
		if tmp.no == id {
			helper.next = tmp.next
			fmt.Printf("Cat=%d\n", id)
		} else {
			fmt.Println("Id not exist...")
		}
	}
	return head
}

func main() {

}
