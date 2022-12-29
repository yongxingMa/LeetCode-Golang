package main

import (
	"fmt"
	"testing"
)

/**
序号：026
标题：重排链表
日期：2022/12/28
*/

func Test026(t *testing.T) {
	s := "abba"
	fmt.Println(countSubstrings(s))
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	slow, fast := head, head
	//找到中间节点
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//temp 后段段链表的反转
	temp := reverseList(slow.Next)
	slow.Next = nil
	mergeList(head, temp)
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}
