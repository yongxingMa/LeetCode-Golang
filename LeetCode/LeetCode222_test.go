package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：222
标题：完全二叉树的节点个数
日期：2023/01/13
*/
func Test222(t *testing.T) {
	var nums = []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}

func countNodes(root *TreeNode) int {
	var getNodeNum func(node *TreeNode) int
	getNodeNum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return getNodeNum(node.Right) + getNodeNum(node.Left) + 1
	}
	return getNodeNum(root)
}
