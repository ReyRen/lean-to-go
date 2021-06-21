package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8001")
	defer conn.Close()

	go func() {
		str := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("os.Stdin.Read err:", err)
				continue
			}
			conn.Write(str[:n])
		}
	}()

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("server is closed!")
			return
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("client read from server: ", string(buf[:n]))
	}
}
