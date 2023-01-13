package LeetCode

/**
序号：110
标题：平衡二叉树
日期：2022/06/12
*/

//自顶向下
//首先计算左右子树的高度，如果左右子树的高度差是否不超过 1，再分别递归地遍历左右子节点，并判断左子树和右子树是否平衡。
//这是一个自顶向下的递归的过程。
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

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
