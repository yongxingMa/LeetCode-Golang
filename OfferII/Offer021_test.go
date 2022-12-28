package main

import (
	"fmt"
	"testing"
)

/**
序号：021
标题：删除链表的倒数第 n 个结点
日期：2022/12/27
*/

func Test021(t *testing.T) {
	s := "abba"
	fmt.Println(countSubstrings(s))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 哨兵指针，防止head节点被删除
	dummy := &ListNode{0, head}
	slow := dummy
	fast := head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
