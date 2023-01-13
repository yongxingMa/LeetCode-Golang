package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：637
标题：二叉树的层平均值
日期：2023/01/12
*/
func Test637(t *testing.T) {
	var nums = []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

func averageOfLevels(root *TreeNode) []float64 {
	ans := []float64{}
	arr := [][]int{}

	depth := 0

	var order func(root *TreeNode, depth int)

	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(arr) == depth {
			arr = append(arr, []int{})
		}
		arr[depth] = append(arr[depth], root.Val)

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	order(root, depth)

	// 遍历数组，求平均值
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := 0; j < len(arr[i]); j++ {
			sum += arr[i][j]
		}
		ans = append(ans, float64(sum)/float64(len(arr[i])))
	}
	return ans

}
