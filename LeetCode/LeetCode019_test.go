package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：19
标题：删除链表倒数第N个节点
日期：2022/06/05
类型：双指针
难度：中等
*/
func Test019(t *testing.T) {
	var s = "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}

/**
在对链表进行操作时，一种常用的技巧是添加一个哑节点（dummy node），它的 \textit{next}next 指针指向链表的头节点。这样一来，我们就不需要对头节点进行特殊的判断了。
*/
//暴力解法
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	length := getLength(head)
//	//哑节点
//	dummy := &ListNode{0, head}
//	cur := dummy
//	for i := 0; i < length-n; i++ {
//		cur = cur.Next
//	}
//	//删除后的拼接
//	cur.Next = cur.Next.Next
//	return dummy.Next
//}
//获取链表的长度
func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

/**
因此我们可以使用两个指针 slow 和 fast 同时对链表进行遍历，并且 设置 超前 n 个节点。当 fast遍历到链表的末尾时，slow 就恰好处于倒数第 n 个节点。
*/
//双指针解决
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//哑节点
	dummy := &ListNode{0, head}
	fast, slow := head, dummy
	//让快指针先走n步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	//当fast=nil时候，slow的下一个就是待删除的节点
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	//删除slow的下一个节点
	slow.Next = slow.Next.Next
	return dummy.Next
}
