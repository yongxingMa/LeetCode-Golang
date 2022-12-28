package LeetCode

/**
序号：145
标题：二叉树的后序遍历
日期：2022/06/12
类型：
*/

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	inorder3(root, &res)
	return res
}
func inorder3(node *TreeNode, res *[]int) {
	if node != nil {
		inorder3(node.Left, res)
		inorder3(node.Right, res)
		*res = append(*res, node.Val)
	}
}
