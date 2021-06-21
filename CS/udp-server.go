package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("net.ResolveUDPAddr err: ", err)
		return
	}

	udpConn, err := net.ListenUDP("udp", srvAddr)
	if err != nil {
		fmt.Println("net.ListenUDP err: ", err)
		return
	}
	defer udpConn.Close()

	buf := make([]byte, 4096)
	n, cltAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("udpConn.ReadFromUDP err: ", err)
		return
	}

	fmt.Printf("server read %v data: ", cltAddr, string(buf[:n]))

	dayTime := time.Now().String()
	_, err = udpConn.WriteToUDP([]byte(dayTime), cltAddr)
	if err != nil {
		fmt.Println("udpConn.WriteToUDP err: ", err)
		return
	}
}
