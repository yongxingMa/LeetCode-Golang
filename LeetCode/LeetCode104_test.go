package LeetCode

/**
序号：104
标题：二叉树的最大深度
日期：2022/06/12
类型：
*/

//深度优先搜索DFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
