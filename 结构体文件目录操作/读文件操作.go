package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("test.file", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("OpenFile err :", err)
		return
	}

	defer f.Close()

	fmt.Println("successful")

	//创建一个带有缓冲区的reader
	reader := bufio.NewReader(f)
	for {
		buf, err := reader.ReadBytes('\n') // 读一行
		if err != nil && err == io.EOF {   // 结束文件标记是要单独读一次的
			fmt.Println("read done")
			return
		} else if err != nil {
			fmt.Println("ReadBytes err: ", err)
		}
		fmt.Println(buf)
	}

}
