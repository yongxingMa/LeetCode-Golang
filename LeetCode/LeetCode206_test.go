package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：206
标题：反转链表
日期：2022/06/01
*/
func Test206(t *testing.T) {
	var height = []int{3, 0, -2, -1, 1, 2}
	fmt.Println(threeSum(height))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		//把第一个节点指向null
		head.Next = prev
		//移动到当前指针
		prev = head
		//把当前的向后移动一下
		head = next
	}
	return prev
}

func reverseList1(head *ListNode) *ListNode {
	//定义新链表
	var prev *ListNode
	for head != nil {
		// 保存next链表
		next := head.Next
		// 当前节点反转指向前边
		head.Next = prev
		//当前节点变成 前驱节点
		prev = head
		// 保存的next节点为当前节点
		head = next
	}
	return prev
}

func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev

}
