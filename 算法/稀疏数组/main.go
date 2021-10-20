package main

import "fmt"

type ValueNode struct {
	row int
	col int
	val int
}

func main() {

	//1 Original Array
	var chessMap [11][11]int
	chessMap[1][2] = 1 // black
	chessMap[2][3] = 2 // blue

	//2 output original array
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//3 convert to sparse array: store to a struct once not zero
	/*var sparseArray []ValueNode
	for i, v := range chessMap {
		for j, v2 := range v {
			valNode := ValueNode{
				row: i,
				col: j,
				val: v2,
			}
			sparseArray := append(sparseArray, valNode)
		}
	}*/

	// 标准的稀疏数组应该还要记录原始数组的规模（行列）
	var sparseArray []ValueNode

	valNode := ValueNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArray = append(sparseArray, valNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				valNode = ValueNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArray = append(sparseArray, valNode)
			}
		}
	}
	for i, valNode := range sparseArray {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}
}
