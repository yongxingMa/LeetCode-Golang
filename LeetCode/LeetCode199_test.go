package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：199
标题：二叉树的右视图
日期：2023/01/12
*/
func Test199(t *testing.T) {
	fmt.Println(climbStairs(8))
}

/**
递归，深度优先搜索
*/
func rightSideView(root *TreeNode) []int {
	arr := make([]int, 0)

	depth := 0

	var order func(root *TreeNode, depth int)

	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		// len(arr)长度，一层只保留一个数，满了就不会添加，先把右节点的加进去
		if len(arr) == depth {
			arr = append(arr, root.Val)
		}
		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}
	order(root, depth)

	return arr

}
