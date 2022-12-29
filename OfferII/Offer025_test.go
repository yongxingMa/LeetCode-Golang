package main

import (
	"fmt"
	"testing"
)

/**
序号：025
标题：链表中的两数相加
日期：2022/12/28
*/

func Test025(t *testing.T) {
	s := "abba"
	fmt.Println(countSubstrings(s))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)
	//初始化节点
	dummy := &ListNode{}
	sumNode := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		remain := (sum + carry) % 10
		carry = (sum + carry) / 10
		sumNode.Next = &ListNode{
			Val: remain,
		}
		sumNode = sumNode.Next
	}
	if carry == 1 {
		sumNode.Next = &ListNode{
			Val: 1,
		}
	}
	return reverseList(dummy.Next)
}
