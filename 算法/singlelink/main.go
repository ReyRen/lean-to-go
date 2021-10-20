package main

import (
	"fmt"
)

type HeroNode struct {
	no       int
	name     string
	nickName string
	next     *HeroNode
}

func InsertHeroNode(head *HeroNode, newNode *HeroNode) {
	tmp := head
	for {
		if tmp.next == nil {
			break
		}
		tmp = tmp.next
	}
	tmp.next = newNode
}

func InsertHeroNodeByIdOrder(head *HeroNode, newNode *HeroNode) {

	tmp := head
	flag := true
	for {
		if tmp.next == nil {
			break
		} else if tmp.next.no >= newNode.no {
			break // tmp==>newNode==>
		} else if tmp.next.no == newNode.no {
			flag = false
			break
		}
		tmp = tmp.next
	}
	if !flag {
		fmt.Println("Already exist NO.=", newNode.no)
	} else {
		newNode.next = tmp.next
		tmp.next = newNode
	}

}

func DeleteNode(head *HeroNode, id int) {
	tmp := head
	flag := false

	for {
		if tmp.next == nil {
			break // empty
		} else if tmp.next.no == id {
			flag = true
			break
		}
		tmp = tmp.next
	}

	if flag {
		tmp.next = tmp.next.next
	} else {
		fmt.Println("Id not exist..")
	}
}

func ListHeroNode(head *HeroNode) {

	tmp := head
	if tmp.next == nil {
		fmt.Println("empty list...")
		return
	}

	for {
		fmt.Printf("[%d, %s, %s]==>", tmp.next.no, tmp.next.name, tmp.next.nickName)
		tmp = tmp.next
		if tmp.next == nil {
			break
		}
	}
}

func main() {
	head := &HeroNode{}

	hero1 := &HeroNode{
		no:       1,
		name:     "songjiang",
		nickName: "jishiyu",
		next:     nil,
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "linchong",
		nickName: "baozitou",
		next:     nil,
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "wuyong",
		nickName: "zhiduoxing",
		next:     nil,
	}

	InsertHeroNode(head, hero3)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero1)
	ListHeroNode(head)
}
