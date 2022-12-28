package main

import (
	"fmt"
	"testing"
)

/**
序号：022
标题：链表中环的入口节点
日期：2022/12/27
*/

func Test022(t *testing.T) {
	s := "abba"
	fmt.Println(countSubstrings(s))
}

// 快慢指针
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for {
		// 环形链表不会存在nil
		if fast == nil || fast.Next == nil {
			return nil
		}
		// 快指针走两步，慢指针走一步
		fast = fast.Next.Next
		slow = slow.Next
		// 表示有环
		if slow == fast {
			break
		}
	}
	// 上面找到相遇的节点以后，重新定义节点
	// 再循环一遍，相遇的地方就是环的入口
	newHead := head
	for {
		if newHead == slow {
			return newHead
		}
		slow = slow.Next
		newHead = newHead.Next
	}
}
