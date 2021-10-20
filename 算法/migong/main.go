package main

import "fmt"

//
//  SetWay
//  @Description:
//  @param myMap
//  @param i test index
//  @param j test index
//
func SetWay(myMap *[8][7]int, i int, j int) bool {
	if myMap[6][5] == 2 {
		return true
	} else {
		if myMap[i][j] == 0 {
			myMap[i][j] = 2
			if SetWay(myMap, i+1, j) {
				// down
				return true
			} else if SetWay(myMap, i, j+1) {
				// right
				return true
			} else if SetWay(myMap, i-1, j) {
				// up
				return true
			} else if SetWay(myMap, i, j-1) {
				// left
				return true
			} else {
				myMap[i][j] = 3
				return false
			}
		} else {
			return false
		}
	}
}

func main() {
	// 2-D array -> map
	var myMap [8][7]int // 1:block 0:none go 2:pass way 3:gone but not pass

	// create wall
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	myMap[3][1] = 1
	myMap[3][2] = 1
	//myMap[1][2] = 1
	//myMap[2][2] = 1
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], "  ")
		}
		fmt.Println()
	}

	SetWay(&myMap, 1, 1)
	fmt.Println("Detect done...")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], "  ")
		}
		fmt.Println()
	}

}
