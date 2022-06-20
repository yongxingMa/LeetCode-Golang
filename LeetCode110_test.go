package main

/**
序号：110
标题：平衡二叉树
日期：2022/06/12
类型：
*/

//自顶向下
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Right) && isBalanced(root.Left)

}

//求二叉树的高度
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
