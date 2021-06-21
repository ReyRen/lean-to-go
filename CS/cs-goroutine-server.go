package main

import (
	"fmt"
	"net"
	"strings"
)

func HandlerConnect(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr()
	fmt.Println(addr, " client in!")

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if "exit\n" == string(buf[:n]) || "exit\r\n" == string(buf[:n]) {
			fmt.Println("server received exit from client")
			return
		}
		if n == 0 {
			fmt.Println("client is closed!")
			return
		}
		if err != nil {
			return
		}
		fmt.Println("server read: ", string(buf[:n]))
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:8001")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go HandlerConnect(conn)
	}

}
