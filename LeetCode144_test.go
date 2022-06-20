package main

/**
序号：094
标题：二叉树的前序遍历
日期：2022/06/12
类型：
*/

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	inorder2(root, &res)
	return res
}
func inorder2(node *TreeNode, res *[]int) {
	if node != nil {
		*res = append(*res, node.Val)
		inorder2(node.Left, res)
		inorder2(node.Right, res)
	}
}
