package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	errFunc(err, "et.Dial err:")
	defer conn.Close()

	httpRequest := "GET /renyuan HTTP/1.1\r\nHost: 127.0.0.1:8000\r\n\r\n"
	conn.Write([]byte(httpRequest))

	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	if n == 0 {
		return
	}
	fmt.Printf("|%s|\n", string(buf[:n]))

}
