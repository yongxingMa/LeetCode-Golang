package LeetCode

/**
序号：559
标题：N 叉树的最大深度
日期：2022/06/12
类型：
*/

type Node3 struct {
	Val      int
	Children []*Node3
}

//深度优先搜索DFS
func maxDepth2(root *Node3) int {
	if root == nil {
		return 0
	}
	maxChildDepth := 0
	//遍历子节点的深度
	for _, child := range root.Children {
		childDepth := maxDepth2(child)
		//当前节点与子节点比较
		if childDepth > maxChildDepth {
			maxChildDepth = childDepth
		}
	}
	return maxChildDepth + 1
}
