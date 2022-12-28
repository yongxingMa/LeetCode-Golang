package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：617
标题：合并二叉树
日期：2022/06/10
类型：遍历/简单
*/
func Test617(t *testing.T) {
	var nums = []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

//定义移动上下左右四种情况
//x+1 y+0 向右移动
//x+0 y+1 向上移动
//x+0 y-1 向下移动
//x-1 y+0 向左移动

//var (
//	dx = []int{1, 0, 0, -1}
//	dy = []int{0, 1, -1, 0}
//)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
可以使用深度优先搜索合并两个二叉树。从根节点开始同时遍历两个二叉树，并将对应的节点进行合并。

两个二叉树的对应节点可能存在以下三种情况，对于每种情况使用不同的合并方式。

如果两个二叉树的对应节点都为空，则合并后的二叉树的对应节点也为空；

如果两个二叉树的对应节点只有一个为空，则合并后的二叉树的对应节点为其中的非空节点；

如果两个二叉树的对应节点都不为空，则合并后的二叉树的对应节点的值为两个二叉树的对应节点的值之和，此时需要显性合并两个节点。

*/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}
