package LeetCode

/**
序号：094
标题：二叉树的中序遍历
日期：2022/06/12
类型：
*/

//递归解法
func inorderTraversal(root *TreeNode) []int {
	var preorder func(*TreeNode)
	// 统计数组
	vals := []int{}
	// 函数定义
	preorder = func(node *TreeNode) {
		//递归终止条件
		if node == nil {
			return
		}
		//遍历左节点
		preorder(node.Left)
		//中序遍历，记录结果
		vals = append(vals, node.Val)
		//遍历右节点
		preorder(node.Right)

	}
	//函数调用
	preorder(root)
	return vals
}

//迭代实现
func inorderTraversal2(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}
