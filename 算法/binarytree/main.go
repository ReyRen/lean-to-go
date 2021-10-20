package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	left  *Hero
	right *Hero
}

// forward: root--left--right
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("No=%d name=%s\n", node.No, node.Name)
		PreOrder(node.left)
		PreOrder(node.right)
	}
}

func main() {
	root := &Hero{
		No:   1,
		Name: "songjiang",
	}
	left := &Hero{
		No:   2,
		Name: "wuyong",
	}
	right1 := &Hero{
		No:   3,
		Name: "lujunyi",
	}
	root.left = left
	root.right = right1
	right2 := &Hero{
		No:   4,
		Name: "linchong",
	}
	right1.right = right2
}
