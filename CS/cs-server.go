package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.listen err")
		return
	}
	defer listener.Close()

	fmt.Println("server waiting client connect...")
	conn, _ := listener.Accept()
	defer conn.Close()
	fmt.Println("server client connect success")

	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	conn.Write(buf)

	fmt.Println("server read from client: ", string(buf[:n]))
}
