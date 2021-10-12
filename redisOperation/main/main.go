package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	_, err = conn.Do("set", "name", "TomJerry")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	/*r, err := conn.Do("get", "name")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	} else {
		//r is interface type
		// r.(string)
	}*/
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	} else {
		fmt.Println(r)
	}

}
