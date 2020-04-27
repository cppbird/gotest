package main

import (
	"fmt"
)

//Num of obj
var Num = 10

func main() {
	// A properity
	var A = []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	// B properity
	var B = []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	fmt.Println(Pac01(A, B, 100))
	fmt.Println(Pac01Gun(A, B, 100))
}

// Pac01 .BV b约束
func Pac01(A, B []int, BV int) int {
	objNum := len(A)
	all := make([][]int, objNum)
	for i := 0; i < objNum; i++ {
		all[i] = make([]int, BV)
	}
	for i := 0; i < objNum; i++ {
		for j := 0; j < BV; j++ {
			if i <= 0 {
				if BV > B[0] {
					all[0][j] = B[0]
				} else {
					all[0][j] = 0
				}
				continue
			}
			if j > A[i] {
				all[i][j] = max(all[i-1][j], all[i-1][j-A[i]]+B[i])
			} else {
				all[i][j] = all[i-1][j]
			}
		}
		fmt.Println(all[i])
	}

	return all[objNum-1][BV-1]
}

// Pac01Gun .BV b约束
func Pac01Gun(A, B []int, BV int) int {
	objNum := len(A)
	all := make([]int, BV)
	for i := 0; i < objNum; i++ {
		for j := BV - 1; j >= 0; j-- {
			if i == 0 {
				if BV > B[0] {
					all[j] = B[0]
				} else {
					all[j] = 0
				}
				continue
			}
			if j > A[i] {
				all[j] = max(all[j], all[j-A[i]]+B[i])
			}
		}
		fmt.Println(all)
	}
	return all[BV-1]
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
