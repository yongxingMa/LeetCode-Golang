package LeetCode

/**
序号：145
标题：二叉树的后序遍历
日期：2022/06/12
类型：
*/

//递归
func postorderTraversal(root *TreeNode) []int {
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
		//遍历右节点
		preorder(node.Right)
		//后序遍历，记录结果
		vals = append(vals, node.Val)
	}
	//函数调用
	preorder(root)
	return vals
}

//迭代
func postorderTraversal2(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}
