package main

import "fmt"

func SelectSort(arr *[5]int) {
	// (*arr)[1] = 600 // = arr[1] = 600
	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
	}
}

func main() {
	arr := [5]int{10, 34, 19, 100, 80}
	SelectSort(&arr)
	fmt.Println(arr)
}
