package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：733
标题：图像渲染
日期：2022/06/09
类型：遍历/简单
*/
func Test733(t *testing.T) {
	var nums = []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

//定义移动上下左右四种情况
//x+1 y+0 向右移动
//x+0 y+1 向上移动
//x+0 y-1 向下移动
//x-1 y+0 向左移动

var (
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

//深度优先算法
//[sr][sc]表示起始的坐标轴
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	//当前像素值
	currColor := image[sr][sc]
	//如果当前的目标一致，直接返回
	//如果不一致，进行深度优先搜索
	if currColor != newColor {
		dfs(image, sr, sc, currColor, newColor)
	}
	return image
}

func dfs(image [][]int, x, y, color, newColor int) {
	//如果等于起始位置color
	if image[x][y] == color {
		//设置为新的color
		image[x][y] = newColor
		for i := 0; i < 4; i++ {
			mx, my := x+dx[i], y+dy[i]
			//大于0 并保证在边界内
			if mx >= 0 && mx < len(image) && my >= 0 && my < len(image[0]) {
				//进行深度优先遍历
				dfs(image, mx, my, color, newColor)
			}
		}
	}
}

//广度优先算法
//先将坐标image[sr][sc]染色作为第一层,然后对其上下左右染色作为第二层;
//再将第二层的分别的上下左右染色,作为第三层,以此类推;

func floodFill2(image [][]int, sr int, sc int, newColor int) [][]int {
	//记录当前的颜色
	currColor := image[sr][sc]
	if currColor == newColor {
		return image
	}
	//len(image) 表示有几行
	// len(image[0]) 表示有几列
	n, m := len(image), len(image[0])
	//定义一个二维数组
	queue := [][]int{}
	//先加入起始坐标值
	queue = append(queue, []int{sr, sc})
	//起始坐标值设置为新颜色
	image[sr][sc] = newColor
	//遍历数组
	//先搜索第一层，再搜索第二层
	for i := 0; i < len(queue); i++ {
		cell := queue[i]
		for j := 0; j < 4; j++ {
			//进行上下左右的遍历
			mx, my := cell[0]+dx[j], cell[1]+dy[j]
			//大于0 并且不出边界 并且等于起始位置
			if mx >= 0 && mx < n && my >= 0 && my < m && image[mx][my] == currColor {
				queue = append(queue, []int{mx, my})
				image[mx][my] = newColor
			}
		}
	}
	return image
}
