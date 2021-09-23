package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	monster := Monster{
		Name:     "zhangsan",
		Age:      12,
		Birthday: "2021-11-11",
		Sal:      800.0,
		Skill:    "niu",
	}

	// serialization
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("ser err...")
	}
	fmt.Printf("ser after=%v", string(data))
}

func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "honghaier"
	a["age"] = 32
	a["address"] = []string{"qqqqq", "wwwwwwwwwwww"}
	//ser
	data, err := json.Marshal(&a)
	if err != nil {
		fmt.Println("ser err...")
	}
	fmt.Printf("ser after=%v", string(data))
}

func testSlice() {
	var slice []map[string]interface{}
	m1 := make(map[string]interface{})
	m1["name"] = "sdasd"
	m1["age"] = 123
	slice = append(slice, m1)
	m2 := make(map[string]interface{})
	m2["name"] = "3333"
	m2["age"] = 121123
	slice = append(slice, m2)
	//ser
	data, err := json.Marshal(&slice)
	if err != nil {
		fmt.Println("ser err...")
	}
	fmt.Printf("ser after=%v", string(data))
}

func main() {
	testStruct()
	testMap()
	testSlice()
}
