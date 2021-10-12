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

	_, err = conn.Do("HMSet", "user02", "name", "James02", "age", 200)
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	r, err := redis.Strings(conn.Do("HMGet", "user02", "name", "age"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	} else {
		fmt.Println(r)
	}

	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}

}
