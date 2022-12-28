package main

import (
	"fmt"
	"testing"
)

/**
序号：013
标题：二维子矩阵的和
日期：2022/12/24
*/

func Test013(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
}

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sums := make([][]int, len(matrix))
	for i, row := range matrix {
		sums[i] = make([]int, len(row)+1)
		for j, v := range row {
			sums[i][j+1] = sums[i][j] + v
		}
	}
	return NumMatrix{sums}
}

func (nm *NumMatrix) SumRegion(row1, col1, row2, col2 int) (sum int) {
	for i := row1; i <= row2; i++ {
		sum += nm.sums[i][col2+1] - nm.sums[i][col1]
	}
	return
}
