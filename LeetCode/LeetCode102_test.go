package LeetCode

/**
序号：102
标题：二叉树层序遍历
日期：2022/01/11
*/

// BFS
func levelOrder6(root *TreeNode) [][]int {
	res := [][]int{}
	// 定义含有跟节点的队列
	arr := []*TreeNode{root}

	if root == nil {
		return res
	}
	for len(arr) > 0 {

		vals := make([]int, len(arr))
		for i := range arr {
			node := arr[0]
			arr = arr[1:]
			vals[i] = node.Val
			if node.Left != nil {
				arr = append(arr, node.Left)
			}
			if node.Right != nil {
				arr = append(arr, node.Right)
			}
		}
		//队列前边的数据出列
		res = append(res, vals)
	}
	return res
}

/*
*
解法一： 二叉树的递归遍历
*/
func levelOrder(root *TreeNode) [][]int {
	arr := [][]int{}

	depth := 0

	var order func(root *TreeNode, depth int)

	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		//拼接空的数组
		if len(arr) == depth {
			arr = append(arr, []int{})
		}
		//拼接节点Val
		arr[depth] = append(arr[depth], root.Val)
		//下一层
		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	order(root, depth)

	return arr
}
