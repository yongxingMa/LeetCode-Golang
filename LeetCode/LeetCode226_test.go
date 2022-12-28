package LeetCode

/**
序号：226
标题：翻转二叉树
日期：2022/06/12
类型：
*/

/**
Homebrew story
*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root

}
