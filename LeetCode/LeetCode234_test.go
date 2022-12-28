package LeetCode

import (
	"testing"
)

/**
序号：234
标题：回文链表
日期：2022/06/12
类型：链表/简单
*/
func Test234(t *testing.T) {
	//var s = "pwwkew"
	//fmt.Println(lengthOfLongestSubstring(s))
}

//type ListNode struct {
//    Val int
//	Next *ListNode
//}

//双指针解法  该方法好像还可以更快?
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//如果是奇数的节点，把中间节点归到左边
	if fast != nil {
		slow = slow.Next
	}
	slow = reverseList(slow)
	fast = head
	for slow != nil {
		if fast.Val != slow.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}
	return true
}

//数组+双指针
func isPalindrome2(head *ListNode) bool {
	vals := []int{}
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	n := len(vals)
	for i, v := range vals[:n/2] {
		if v != vals[n-1-i] {
			return false
		}
	}
	return true
}
