package LeetCode

import "math"

/**
序号：098
标题：验证二叉搜索树
日期：2023/02/02
*/

// 二叉搜索树 左子树所有节点小于中间节点，右子树所有节点大于中间节点
func isValidBST(root *TreeNode) bool {
	// 二叉搜索树也可以是空树
	if root == nil {
		return true
	}
	// 由题目中的数据限制可以得出min和max
	return checkBST(root, math.MinInt64, math.MaxInt64)
}
func checkBST(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}

	if min >= int64(node.Val) || max <= int64(node.Val) {
		return false
	}
	// 分别对左子树和右子树递归判断，如果左子树和右子树都符合则返回true
	return checkBST(node.Right, int64(node.Val), max) && checkBST(node.Left, min, int64(node.Val))
}
