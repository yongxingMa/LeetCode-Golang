package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：203
标题：移除链表元素
日期：2022/12/30
*/
func Test203(t *testing.T) {
	var height = []int{3, 0, -2, -1, 1, 2}
	fmt.Println(threeSum(height))
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return dummy.Next
}
