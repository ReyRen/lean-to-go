package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f_r, err := os.Open("srcFile")
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer f_r.Close()

	//创建写文件
	f_w, err := os.Create("dstFile")
	if err != nil {
		fmt.Println("create err:", err)
		return
	}
	defer f_w.Close()

	//从读文件中获取数据，放到缓冲区中
	buf := make([]byte, 4096)
	for {
		n, err := f_r.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("finished")
			return
		}
		f_w.Write(buf[:n]) // 读多少写多少
	}
}
