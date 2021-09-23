package testCase2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {
	//serialize
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err =", err)
		return false
	}
	//save to file
	filePath := "./monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("save file err =", err)
		return false
	}
	return true
}
func (this *Monster) reStore() bool {
	filePath := "./monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file err =", err)
		return false
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("unmarshal err =", err)
		return false
	}
	return true
}
