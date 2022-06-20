package main

/**
序号：022
标题：链表中倒数第k个节点
日期：2022/06/12
类型：双指针
难度：简单
*/

//哈希表记录法

//快慢指针解法
/**
倒数第k个 就是正数的第n-k+1个
*/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	//快指针先走k-1次
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
