package main

import (
	"fmt"
	"os"
)

func main() {

	//获取用户输入的目录路径
	fmt.Println("请输入查询的目录：")
	var path string
	fmt.Scan(&path)

	//打开目录
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("openfile err:", err)
		return
	}
	defer f.Close()

	// 读取目录项
	info, err := f.Readdir(-1) // -1表示读取所有目录项
	//遍历返回的切片
	for _, fileInfo := range info {
		if fileInfo.IsDir() {
			fmt.Println(fileInfo.Name(), "是一个目录")
		} else {
			fmt.Println(fileInfo.Name(), "是一个文件")

		}
	}
}
