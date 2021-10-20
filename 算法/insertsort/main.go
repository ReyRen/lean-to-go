package main

import "fmt"

func InsertSort(arr *[5]int) {

	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1

		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		// insert
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}
}

func main() {
	arr := [5]int{23, 0, 12, 56, 34}
	InsertSort(&arr)
	fmt.Println(arr)
}
