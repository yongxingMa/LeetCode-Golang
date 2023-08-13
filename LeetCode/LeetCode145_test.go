package LeetCode

/**
序号：145
标题：二叉树的后序遍历
日期：2022/06/12
类型：
*/

// 递归
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

// 迭代
func postorderTraversal2(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	var prev *TreeNode
	// root不为空或者栈里有数据 继续处理
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			//先寻找左子树
			root = root.Left
		}
		//当没有左子树的时候，弹出栈顶的节点
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		// 如果当前节点有右子树或栈顶的节点的右子树已经遍历过，记录当前节点的值，并标记为先前节点，清空root，继续下一轮弹出节点
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			//如果有右子树，继续入栈，因为后边可能还有左子树
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}
