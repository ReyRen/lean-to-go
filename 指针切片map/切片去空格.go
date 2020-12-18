package main

import "fmt"

func noEmpty(data []string) []string {
	out := data[:0] // == make([]string, 0)
	for _, str := range data {
		if str != "" {
			out = append(out, str)
		}
	}
	return out
}

//直接在原字符串上操作
func noEmpty2(data []string) []string {
	i := 0
	for _, str := range data {
		if str != "" {
			data[i] = str
			i++
		}
	}
	return data[:i]
}

func main() {
	data := []string{"red", "", "black", "", "", "pink"}
	afterData := noEmpty(data)
	fmt.Println(afterData)
}
