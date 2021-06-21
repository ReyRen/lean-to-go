package main

import (
	"fmt"
	"net"
	"os"
)

func recvFile(conn net.Conn, fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer f.Close()

	// read from tcp to local
	buf := make([]byte, 4096)
	for true {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Println("receive file done")
		}
		f.Write(buf[:n])
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	fileName := string(buf[:n])
	_, err = conn.Write([]byte("ok"))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}

	recvFile(conn, fileName)
}
