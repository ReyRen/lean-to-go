package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	//fmt.Println(list) // /tmp/___go_build_fileAttribute_go a b c
	if len(list) != 2 {
		fmt.Println("format is \"go run yy.go xx.file\"")
		return
	}

	path := list[1]
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err:", err)
		return
	}
	fmt.Println("filename: ", fileInfo.Name())
	fmt.Println("filesize: ", fileInfo.Size())

}
