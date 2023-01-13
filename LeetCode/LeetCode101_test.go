package LeetCode

/**
序号：101
标题：对称二叉树
日期：2023/01/13
*/

//递归解法
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(left, right *TreeNode) bool {
	// 左子树为空，右子树为空
	if left == nil && right == nil {
		return true
	}
	// 其中一个为空
	if left == nil || right == nil {
		return false
	}
	// 左子树和右子树值不相等
	if left.Val != right.Val {
		return false
	}
	return check(left.Left, right.Right) && check(left.Right, right.Left)
}
