package LeetCode

import (
	"testing"
)

/**
序号：021
标题：合并两个有序链表
日期：2022/06/11
类型：递归回溯/中等
*/
func Test024(t *testing.T) {
	//var s = "pwwkew"
	//fmt.Println(lengthOfLongestSubstring(s))
}

//type ListNode struct {
//    Val int
//	Next *ListNode
//}

//递归解法
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next
		//pre=list[(i+2)-1]
		pre = head
		//head=list[(i+2)]
		head = next
	}
	return dummy.Next
}
