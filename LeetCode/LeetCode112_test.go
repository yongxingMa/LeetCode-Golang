package LeetCode

/**
序号：112
标题：二叉树路径总和
日期：2023/01/17
*/

//递归+回溯
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val // 将targetSum在遍历每层的时候都减去本层节点的值
	// 递归终止条件
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}
	// 否则递归找
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
