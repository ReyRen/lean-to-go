package main

import (
	"fmt"
	"strings"
)

func makeSuffic(suffix string) func(string) string {

	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

//这里是suffix和返回的func(name string)string{}构成了闭包。
//闭包可以形象的比喻成宇宙飞船中包含的小的宇宙飞机。它们在宇宙飞船中打造生成发射出去用于外界探索，但是它们共享这宇宙飞船中的一切。

func main() {
	f := makeSuffic(".jpg")
	fmt.Println(f("xxx"))
	fmt.Println(f("yy.jpg"))
}
