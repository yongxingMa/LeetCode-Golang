package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：054
标题：螺旋矩阵
日期：2023/07/31
*/
func Test054(t *testing.T) {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	res := []int{}
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for top < bottom && left < right {
		for i := left; i < right; i++ {
			res = append(res, matrix[top][i])
		}
		for i := top; i < bottom; i++ {
			res = append(res, matrix[i][right])
		}
		for i := right; i > left; i-- {
			res = append(res, matrix[bottom][i])
		}
		for i := bottom; i > top; i-- {
			res = append(res, matrix[i][left])
		}
		right--
		top++
		bottom--
		left++
	}
	if top == bottom {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
	} else if left == right {
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][left])
		}
	}
	return res
}
