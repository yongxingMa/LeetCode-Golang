package main

import (
	"fmt"
	"math"
	"testing"
)

/**
序号：994
标题：腐烂的橘子
日期：2022/06/10
类型：遍历/中等
*/
func Test994(t *testing.T) {
	var nums = []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

//定义移动上下左右四种情况
//x+1 y+0 向右移动
//x+0 y+1 向上移动
//x+0 y-1 向下移动
//x-1 y+0 向左移动

//var (
//	dx = []int{1, 0, 0, -1}
//	dy = []int{0, 1, -1, 0}
//)

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// BFS
func orangesRotting(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	news := make([][]int, n)
	list := make([][]int, 0)
	for r := range grid {
		news[r] = make([]int, m)
		for rr := range grid[r] {
			switch grid[r][rr] {
			case 0:
				news[r][rr] = math.MaxInt32 - 1
			case 1:
				news[r][rr] = math.MaxInt32
			case 2:
				news[r][rr] = 0
				list = append(list, []int{r, rr})
			}
		}
	}

	for len(list) != 0 {
		colum := list[0]
		list = list[1:]
		for r := range dx {
			mx, my := colum[0]+dx[r], colum[1]+dy[r]
			if mx >= 0 && mx < n && my >= 0 && my < m && news[mx][my] == math.MaxInt32 {
				if news[colum[0]][colum[1]] < math.MaxInt32-2 {
					news[mx][my] = news[colum[0]][colum[1]] + 1
				}
				list = append(list, []int{mx, my})
			}
		}
	}
	var max int
	for r := range news {
		for rr := range news[r] {
			if news[r][rr] == math.MaxInt32 {
				return -1
			}
			if news[r][rr] > math.MaxInt32-2 {
				news[r][rr] = 0
			} else if news[r][rr] > max {
				max = news[r][rr]
			}
		}
	}
	return max
}
