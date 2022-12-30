package main

import (
	"fmt"
	"testing"
)

/**
序号：027
标题：回文链表
日期：2022/12/29
*/

func Test027(t *testing.T) {
	s := "abba"
	fmt.Println(countSubstrings(s))
}

func isPalindrome(head *ListNode) bool {

	rehead := reverseList(head)
	for head != nil {
		if head.Val == rehead.Val {
			fmt.Println(head.Val + rehead.Val)
			head = head.Next
			rehead = rehead.Next
		} else {
			return false
		}

	}
	return false
}
