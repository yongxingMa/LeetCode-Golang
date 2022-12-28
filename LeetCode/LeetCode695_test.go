package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：695
标题：岛屿的最大面积
日期：2022/06/09
类型：遍历/中等
*/
func Test695(t *testing.T) {
	var grid = [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println(maxAreaOfIsland(grid))
}

//找出所有岛屿的面积，输出最大值的面积

//定义移动上下左右四种情况
//x+1 y+0 向右移动
//x+0 y+1 向上移动
//x+0 y-1 向下移动
//x-1 y+0 向左移动

//var (
//	dx = {}int{1, 0, 0, -1}
//	dy = {}int{0, 1, -1, 0}
//)

//深度优先算法
//{sr}{sc}表示起始的坐标轴
func maxAreaOfIsland(grid [][]int) int {
	//待返回的结果
	res := 0
	//先拿出来一行，代表第i行，row代表一个数组
	for i, row := range grid {
		//表示第i行的的第一列，j去遍历，col表示具体的数
		for j, col := range row {
			if col == 0 {
				continue
			}
			//判断是否是岛屿
			area := areaIsland(grid, i, j)
			//如果有更大的岛屿，更新数值
			if area > res {
				res = area
			}
		}
	}
	return res
}
func areaIsland(grid [][]int, x, y int) int {
	//如果在边界内，并且不是水
	if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] != 0 {
		//访问过的土地设置为0
		grid[x][y] = 0
		total := 1
		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			total += areaIsland(grid, nx, ny)
		}
		return total
	} else {
		return 0
	}
}
