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

	_, err = conn.Do("HSet", "user01", "name", "James")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	r, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	} else {
		fmt.Println(r)
	}
	r2, err := redis.Int(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	} else {
		fmt.Println(r2)
	}

}
