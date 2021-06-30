package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func Manager() {
	onlineMap = make(map[string]Client)

	for true {
		// listen global channel
		msg := <-message

		for _, v := range onlineMap {
			v.C <- msg
		}
	}
}

// user struct
type Client struct {
	C    chan string
	Name string
	Addr string
}

// map online client
var onlineMap map[string]Client

// globale channel
var message = make(chan string)

func WritreMsgToClient(clnt Client, conn net.Conn) {
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(clnt Client, msg string) (buf string) {
	bufs := "[" + clnt.Addr + "] " + clnt.Name + ":" + buf
	return bufs
}

func HandlerConnect(conn net.Conn) {

	defer conn.Close()
	// user active channel
	hasData := make(chan bool)

	// get client ip port
	netAddr := conn.RemoteAddr().String()

	// create new client struct
	clnt := Client{
		make(chan string),
		netAddr,
		netAddr,
	}

	onlineMap[netAddr] = clnt

	go WritreMsgToClient(clnt, conn)

	// online msg
	//message <- "[" + netAddr + "] " + clnt.Name + " login"
	message <- MakeMsg(clnt, "login")

	//exit channel
	isQuit := make(chan bool)

	//handle client sent msg
	go func() {
		buf := make([]byte, 4096)
		for true {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Printf("client %s closed\n", clnt.Name)
				return
			}
			if err != nil {
				fmt.Println("conn.Read err", err)
				return
			}
			//broadcast msg
			msg := string(buf[:n-1]) // remove /r /r/n
			if msg == "who" && len(msg) == 3 {
				conn.Write([]byte("online user list:\n"))
				for _, user := range onlineMap {
					userInfo := user.Addr + ":" + user.Name + "\n"
					conn.Write([]byte(userInfo))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				newName := strings.Split(msg, "|")[1] // = msg[8:]
				clnt.Name = newName
				onlineMap[netAddr] = clnt
				conn.Write([]byte("rename successful\n"))
			} else {
				message <- MakeMsg(clnt, msg)
			}

			// read is block func, so stand here would means activate
			hasData <- true
		}
	}()

	for true {
		select {
		case <-isQuit:
			/*
				这里需要close一下clinet的channel从而能将WritreMsgToClient这个goroutine结束
				因为，在go程中的子go程，当父go程退出时，并不会影响子go程的运行，但是父进程退出
				那其下的所有go程都结束了。但父进程中执行goexit，也不会影响到子go程，因为goexit只是
				go程退出函数。
			*/
			close(clnt.C)
			delete(onlineMap, clnt.Addr)
			message <- MakeMsg(clnt, "logout")
			return // exit all self goroutine
		case <-hasData:
			// nothing to do. Aim to reset timeAfter timer
		case <-time.After(time.Second * 10):
			close(clnt.C)
			delete(onlineMap, clnt.Addr)
			message <- MakeMsg(clnt, "logout")
			return // exit all self goroutine
		}
	}
}

func main() {

	// socket listen
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	// create manager goroutine
	go Manager()

	// accept
	for true {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}
		go HandlerConnect(conn)
	}

}
