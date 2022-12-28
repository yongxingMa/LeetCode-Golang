package LeetCode

/**
序号：876
标题：链表的中间节点
日期：2022/06/05
类型：双指针
难度：简单
*/

//快慢指针解法
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	//当快指针到了以后，返回慢指针的位置
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
