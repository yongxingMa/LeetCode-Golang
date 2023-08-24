package LeetCode

/**
序号：103
标题：二叉树的锯齿形层序遍历
之字型输出
日期：2022/01/11
*/

// BFS
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	// 奇偶交替，取反
	for even := false; len(q) > 0; even = !even {
		n := len(q)
		vals := make([]int, n) // 大小已知
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			// 如果是偶数
			if even {
				vals[n-1-i] = node.Val // 倒着添加
			} else {
				// 奇数 直接添加
				vals[i] = node.Val
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, vals)
	}
	return ans
}
