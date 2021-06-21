package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(conn net.Conn, filePath string) {
	// readonly
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open err:", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 4096)
	for true {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("send file completely")
			} else {
				fmt.Println("f.Read err:", err)
			}
			return
		}
		conn.Write(buf[:n])

	}
}

func main() {

	list := os.Args
	if len(list) != 2 {
		fmt.Println("FORMAT: go run xxx.go FILEPATH")
		return
	}

	filePath := list[1]
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("os.Stat err:", err)
		return
	}
	fileName := fileInfo.Name()

	conn, err := net.Dial("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	// send file name to receiver
	_, err = conn.Write([]byte(fileName))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
	// read respond from receiver
	buf := make([]byte, 16)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	if "ok" == string(buf[:n]) {
		sendFile(conn, filePath)
	}

}
