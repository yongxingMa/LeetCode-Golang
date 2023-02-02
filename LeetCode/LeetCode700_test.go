package LeetCode

/**
序号：700
标题：二叉搜索树中的搜索
日期：2023/02/02
*/

func searchBST(root *TreeNode, val int) *TreeNode {
	// 如果root为空，或者找到这个数值了，就返回root节点。
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}
