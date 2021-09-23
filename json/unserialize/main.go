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

func unMarshalStruct() {
	str := "{\"Age\":500,\"Birthday\":\"1992-1-1\",\"Name\":\"zhangsan\",\"Sal\":50.0,\"Skill\":\"sharen\"}"
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster) //这里必须要用引用，才能改变monster的值
	if err != nil {
		fmt.Println("unmarshal err...")
	}

	fmt.Println(monster)
}

func unMarshalMap() {
	str := "{\"address\":\"111111\",\"age\":30}"

	var a map[string]interface{} // no need to make(make already encapulat in unmarshal)
	json.Unmarshal([]byte(str), &a)
}

func main() {
	unMarshalStruct()
}
