package main

import (
	"fmt"
	"lean-to-go/chatRoom-MVC/server/model"
	"net"
	"time"
)

func processMain(conn net.Conn) {

	defer conn.Close()
	processer := &Processer{Conn: conn}
	err := processer.processMainWorker()
	if err != nil {
		fmt.Println("processer.processMainWorker err=", err)
		return
	}
}

// initialize UserDao
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {

	initPool("127.0.0.1:6379", 8, 0, 300*time.Second)
	initUserDao()

	fmt.Println("Server is open 8889")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("Waiting for the client...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		go processMain(conn)
	}

}
