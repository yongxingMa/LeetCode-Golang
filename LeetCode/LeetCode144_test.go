package LeetCode

/**
序号：094
标题：二叉树的前序遍历
日期：2022/06/12
类型：
*/

//递归
func preorderTraversal(root *TreeNode) []int {
	var preorder func(*TreeNode)
	// 统计数组
	vals := []int{}
	// 函数定义
	preorder = func(node *TreeNode) {
		//递归终止条件
		if node == nil {
			return
		}
		//前序遍历，第一个值放入数组
		vals = append(vals, node.Val)
		//遍历左节点
		preorder(node.Left)
		//遍历右节点
		preorder(node.Right)
	}
	//函数调用
	preorder(root)
	return vals
}

//迭代
func preorderTraversal2(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}
