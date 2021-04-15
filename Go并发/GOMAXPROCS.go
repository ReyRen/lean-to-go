package main

import (
	"fmt"
	"runtime"
)

func main() {
	_ = runtime.GOMAXPROCS(1) // single cpu
	// n return the last CPU num setted
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
