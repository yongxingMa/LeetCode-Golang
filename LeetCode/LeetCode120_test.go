package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：120
标题：三角形最小路径和
日期：2022/06/15
类型：
*/
func Test120(t *testing.T) {
	fmt.Println(climbStairs(8))
}

/**
动态规划可以用递归或迭代的方法去实现
*/
func minimumTotal(triangle [][]int) int {
	if triangle == nil {
		return 0
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
