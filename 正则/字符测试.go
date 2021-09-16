package main

import (
	"fmt"
	"regexp"
)

func main() {

	str := "ac a6bc a7c mfc cat 8ca azc cba"
	// parse and compile regexp
	//ret := regexp.MustCompile(`a.c`) // `` means original string
	ret := regexp.MustCompile(`a[0-9]c`) // `` means original string
	alls := ret.FindAllStringSubmatch(str, -1)
	fmt.Println(alls)
}
