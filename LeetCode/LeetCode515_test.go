package LeetCode

import (
	"fmt"
	"math"
	"testing"
)

/**
序号：515
标题：在每个树行中找最大值
日期：2023/01/12
*/
func Test515(t *testing.T) {
	var nums = []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

func largestValues(root *TreeNode) []int {
	ans := []int{}
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

	// 遍历数组，求每一层的最大值
	for i := 0; i < len(arr); i++ {
		maxNum := math.MinInt32
		for j := 0; j < len(arr[i]); j++ {
			maxNum = max(maxNum, arr[i][j])
		}
		ans = append(ans, maxNum)
	}
	return ans

}
