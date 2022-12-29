package main

import (
	"fmt"
	"testing"
)

/**
序号：024
标题：反转链表
日期：2022/12/27
*/

func Test024(t *testing.T) {
	s := "abba"
	fmt.Println(countSubstrings(s))
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		// 存储
		next := head.Next
		// 指向前一个节点
		head.Next = prev
		//prev往右走一步
		prev = head
		// 赋值，与第一步的存储对应
		head = next
	}
	return prev
}
