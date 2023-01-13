package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：404
标题：左叶子之和
日期：2023/01/13
*/
func Test404(t *testing.T) {
	var nums = []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftValue := sumOfLeftLeaves(root.Left) // 左

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftValue = root.Left.Val // 中
	}

	rightValue := sumOfLeftLeaves(root.Right) // 右

	return leftValue + rightValue
}
