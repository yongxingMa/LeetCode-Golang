package LeetCode

/**
序号：094
标题：二叉树的中序遍历
日期：2022/06/12
类型：
*/

//递归解法
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	inorder(root, &res)
	return res
}
func inorder(node *TreeNode, res *[]int) {
	if node != nil {
		inorder(node.Left, res)
		*res = append(*res, node.Val)
		inorder(node.Right, res)
	}
}

//迭代用栈实现
//func inorderTraversal2(root *TreeNode) (res []int) {
//	stack := []*TreeNode{}
//	for root != nil || len(stack) > 0 {
//		//节点不为空，先入栈
//		for root != nil {
//			stack = append(stack, root)
//			//寻找左子树
//			root = root.Left
//		}
//		//栈顶弹出来
//		root = stack[len(stack)-1]
//		//删掉栈顶的数
//		stack = stack[:len(stack)-1]
//		res = append(res, root.Val)
//		root = root.Right
//	}
//	return
//}
