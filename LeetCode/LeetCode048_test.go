package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：48
标题：旋转图像
日期：2023/08/14

	matrix = [[1,2,3],[4,5,6],[7,8,9]]
	>> [[7,4,1],[8,5,2],[9,6,3]]
*/
func Test048(t *testing.T) {
	var height = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(height, 0))
}

// 临时数组
// 第一行的第 x 个元素在旋转后恰好是倒数第一列的第 x 个元素。
func rotate2(matrix [][]int) {
	n := len(matrix)
	//临时数组
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}
	for i, row := range matrix {
		for j, v := range row {
			tmp[j][n-1-i] = v
		}
	}
	copy(matrix, tmp)
}
