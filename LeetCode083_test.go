package main

import (
	"testing"
)

/**
序号：083
标题：删除排序链表中的重复元素
日期：2022/06/11
类型：递归回溯/中等
*/
func Test083(t *testing.T) {
	//var s = "pwwkew"
	//fmt.Println(lengthOfLongestSubstring(s))
}

//type ListNode struct {
//    Val int
//	Next *ListNode
//}

//递归解法
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curNode := head
	for curNode.Next != nil {
		//当前值和下一个值相等，把下一个值删掉
		if curNode.Val == curNode.Next.Val {
			curNode.Next = curNode.Next.Next
		} else {
			//往下移动
			curNode = curNode.Next
		}
	}
	return head
}
