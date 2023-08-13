package LeetCode

/**
序号：662
标题：二叉树最大宽度
日期：2023/07/30
*/

type pair struct {
	node  *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	ans := 1
	q := []pair{{root, 1}}
	for q != nil {
		//最右边的节点值减去最左边的节点值
		ans = max(ans, q[len(q)-1].index-q[0].index+1)
		// 保存临时数据
		tmp := q
		//每次清空，只遍历当前层的数据
		q = nil
		for _, p := range tmp {
			if p.node.Left != nil {
				q = append(q, pair{p.node.Left, p.index * 2})
			}
			if p.node.Right != nil {
				q = append(q, pair{p.node.Right, p.index*2 + 1})
			}
		}
	}
	return ans
}

func widthOfBinaryTree2(root *TreeNode) int {
	ans := 1
	newroot := []pair{{root, 1}}
	for newroot != nil {
		ans = max(ans, newroot[len(newroot)-1].index-newroot[0].index+1)
		tmp := newroot
		newroot = nil
		for _, t := range tmp {
			if t.node.Left != nil {
				newroot = append(newroot, pair{t.node.Left, t.index * 2})
			}
			if t.node.Right != nil {
				newroot = append(newroot, pair{t.node.Right, t.index*2 + 1})
			}
		}
	}
	return ans
}
