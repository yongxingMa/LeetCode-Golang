package main

import (
	"testing"
)

/**
序号：141
标题：环形链表
日期：2022/06/12
类型：链表/简单
*/
func Test141(t *testing.T) {
	//var s = "pwwkew"
	//fmt.Println(lengthOfLongestSubstring(s))
}

//type ListNode struct {
//    Val int
//	Next *ListNode
//}

//使用快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	//环形链表是没有nil的
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
