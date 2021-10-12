package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//global pool
var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
		TestOnBorrow:    nil,
		MaxIdle:         8,
		MaxActive:       0,
		IdleTimeout:     100,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "tttttttttt")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("redis.String err=", err)
		return
	} else {
		fmt.Println(r)
	}

}
