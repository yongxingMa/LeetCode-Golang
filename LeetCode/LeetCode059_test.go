package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：059
标题：螺旋矩阵 II
日期：2022/12/29
*/
func Test059(t *testing.T) {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1
	tar := n * n
	// 二维数组初始化
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for num <= tar {
		//记录 123
		//记录 9
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		//记录 45
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		//记录 67
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		//记录 8
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}
