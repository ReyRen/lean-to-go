package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	// 1. 注册回调函数，该函数会在服务器被访问时，自动被调用
	http.HandleFunc("/renyuan", handler)

	//绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8008", nil) //第二个函数也是个回调函数，表示服务器监听接受到访问的时候，应对连接的函数，一般为nil表示，走net/http自行提供的
}
