package LeetCode

/**
序号：104
标题：二叉树的最大深度
日期：2022/06/12
类型：
*/

// 深度优先搜索DFS
// 后序遍历，先求它的左子树的深度，再求右子树的深度，最后取左右深度最大的数值 再 + 1
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
