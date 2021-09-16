package main

import (
	"fmt"
	"net"
	"os"
)

func errFunc(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		os.Exit(1) // terminate current process
	}
}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8008")
	errFunc(err, "net.Listen err:")
	defer listener.Close()

	conn, err := listener.Accept()
	errFunc(err, "listener.Accept err:")
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if n == 0 {
		return
	}
	errFunc(err, "conn.Read err:")

	fmt.Printf("|%s|\n", string(buf[:n]))
}
