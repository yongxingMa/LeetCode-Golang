package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：429
标题：N 叉树的层序遍历
日期：2023/01/13
*/
func Test429(t *testing.T) {
	var nums = []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}

type Node struct {
	Val      int
	Children []*Node
}

// BFS广度优先搜索
func levelOrder5(root *Node) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	//模拟队列
	var queue []*Node
	queue = append(queue, root)
	for len(queue) != 0 {
		var level []int
		size := len(queue)
		for i := 0; i < size; i++ {
			root := queue[0]
			//队列弹出一个
			queue = queue[1:]
			//队列入队子节点
			queue = append(queue, root.Children...)
			//记录每一层的数据
			level = append(level, root.Val)
		}
		//拼接每一层数据，记录返回结果
		ret = append(ret, level)
	}

	return ret
}
