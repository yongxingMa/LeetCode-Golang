package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：513
标题：找树左下角的值
日期：2023/01/17
*/
func Test513(t *testing.T) {
	var nums = []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

/**
  使用广度优先搜索遍历每一层的节点。
  在遍历一个节点时，需要先把它的非空右子节点放入队列，然后再把它的非空左子节点放入队列，这样才能保证从右到左遍历每一层的节点。
  广度优先搜索所遍历的最后一个节点的值就是最底层最左边节点的值。
*/
func findBottomLeftValue(root *TreeNode) int {
	node := root
	q := []*TreeNode{root}
	for len(q) > 0 {
		node = q[0]
		q = q[1:]
		//放入非空右子节点
		if node.Right != nil {
			q = append(q, node.Right)
		}
		//放入非空左子节点
		if node.Left != nil {
			q = append(q, node.Left)
		}
	}
	return node.Val
}
