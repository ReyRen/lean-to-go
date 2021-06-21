package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	defer conn.Close()

	conn.Write([]byte("Are you server?"))

	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	fmt.Println("client read from server: ", string(buf[:n]))
}
