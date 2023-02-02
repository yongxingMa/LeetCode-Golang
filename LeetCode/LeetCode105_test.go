package LeetCode

/**
序号：105
标题：从前序与中序遍历序列构造二叉树
日期：2023/02/02
类型：
*/

func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}
	left := findRootIndex2(preorder[0], inorder)
	root := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:left+1], inorder[:left]),
		Right: buildTree(preorder[left+1:], inorder[left+1:])}
	return root
}
func findRootIndex2(target int, inorder []int) int {
	for i := 0; i < len(inorder); i++ {
		if target == inorder[i] {
			return i
		}
	}
	return -1
}
