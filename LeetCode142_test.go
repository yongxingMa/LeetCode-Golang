package main

import (
	"testing"
)

/**
序号：142
标题：环形链表 II
日期：2022/06/12
类型：链表/中等
*/
func Test142(t *testing.T) {
	//var s = "pwwkew"
	//fmt.Println(lengthOfLongestSubstring(s))
}

//type ListNode struct {
//    Val int
//	Next *ListNode
//}

//使用hash表

//使用快慢指针
func detectCycle(head *ListNode) *ListNode {
	isloop := false
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	//环形链表是没有nil的
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			isloop = true
			break
		}
	}
	if isloop {
		//慢指针移动到头节点
		slow = head
		for slow != fast {
			fast = fast.Next
			slow = slow.Next
		}
		return slow
	}
	return nil
}
