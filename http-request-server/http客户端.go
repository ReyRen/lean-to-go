package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http.Get err:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Header:", resp.Header)
	fmt.Println("Status:", resp.Status)
	fmt.Println("StatusCode:", resp.StatusCode)
	fmt.Println("Proto:", resp.Proto)

	buf := make([]byte, 4096)
	var result string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 { // = io.EOF
			fmt.Println("read finish")
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("resp.Body.Read err:", err)
			break
		}
		result += string(buf[:n])
	}
	fmt.Printf("|%v|\n", result)
	return
}
