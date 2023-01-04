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

//解法一 额外数组
func isPalindrome3(head *ListNode) bool {
	// 头节点为空，则为回文链表
	if head == nil {
		return true
	}
	// 遍历链表，加入数组
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	// 遍历数组，前后分别对称判断
	n := len(nums)
	for i := 0; i < n; i++ {
		if i <= n-1-i && nums[i] != nums[n-1-i] {
			return false
		}
	}
	return true
}

//解法二：快慢指针
func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// slow.Next是下半个链表
	// 把后半部分进行反转，这里不能对整个链表进行反转
	secondHalfStart := reverseList(slow.Next)

	// 判断是否回文
	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	// 还原链表并返回结果（可选）
	// slow.Next = reverseList(secondHalfStart)
	return result
}
