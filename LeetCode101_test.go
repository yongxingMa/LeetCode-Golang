package main

/**
序号：101
标题：对称二叉树
日期：2022/06/12
类型：
*/

//递归解法
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
}
