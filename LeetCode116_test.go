package main

import (
	"fmt"
	"testing"
)

/**
序号：116
标题：填充每个节点的下一个右侧节点指针
日期：2022/06/10
类型：遍历/中等
*/
func Test116(t *testing.T) {
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

//完美二叉树
//type Node struct {
//	val   int
//	Left  Node
//	Right Node
//	Next  Node
//}

/**
将二叉树的每一层节点都连接起来形成一个链表，因此直观的做法我们可以对二叉树进行层次遍历，在层次遍历的过程中将我们将二叉树每一层的节点拿出来遍历并连接。
层次遍历基于广度优先搜索，它与广度优先搜索的不同之处在于，广度优先搜索每次只会取出一个节点来拓展，而层次遍历会每次将队列中的所有元素都拿出来拓展，
这样能保证每次从队列中拿出来遍历的元素都是属于同一层的，因此我们可以在遍历的过程中修改每个节点的 next 指针，同时拓展下一层的新队列。

*/
//func connect(root *Node) *Node {
//	if root == nil {
//		return nil
//	}
//	// 初始化队列同时将第一层节点加入队列中，即根节点
//	queue := []*Node{root}
//	for len(queue) > 0 {
//		tmp := queue
//		//设置为空，如果下面有子节点，会继续添加进来，继续走循环
//		queue = nil
//		// 遍历这一层的所有节点
//		for i, node := range tmp {
//			// 连接
//			// 当前层如果没有越界，指向右侧的接待你
//			if i+1 < len(tmp) {
//				node.Next = tmp[i+1]
//			}
//			// 拓展下一层节点
//			if node.Left != nil {
//				queue = append(queue, node.Left)
//			}
//			if node.Right != nil {
//				queue = append(queue, node.Right)
//			}
//		}
//	}
//	// 返回根节点
//	return root
//}
