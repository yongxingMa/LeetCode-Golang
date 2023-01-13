package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：117
标题：填充每个节点的下一个右侧节点指针 II
日期：2023/01/13
*/
func Test117(t *testing.T) {
	var nums = []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

//解法二 层序遍历
//初始状态下，所有 next 指针都被设置为 NULL,所以不需要考虑右节点的情况
func connect3(root *Node2) *Node2 {
	if root == nil {
		return root
	}
	// 初始化队列同时将第一层节点加入队列中，即根节点
	queue := []*Node2{root}

	// 循环迭代的是层数
	for len(queue) > 0 {
		tmp := queue
		queue = nil

		// 遍历这一层的所有节点
		for i, node := range tmp {
			// 连接
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}

			// 拓展下一层节点
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	// 返回根节点
	return root
}
