package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：002
标题：两数相加
日期：2022/06/11
类型：/简单
*/
func Test002(t *testing.T) {
	var nums = []int{3, 3}
	fmt.Println(twoSum001(nums, 6))
}

/**
小技巧：对于链表问题，返回结果为头结点时，通常需要先初始化一个预先指针 pre，该指针的下一个节点指向真正的头结点head。
使用预先指针的目的在于链表初始化时无可用节点值，而且链表构造过程需要指针移动，进而会导致头指针丢失，无法返回结果。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//定义一个新链表伪指针，用来指向头指针，返回结果
	var prev *ListNode
	var head *ListNode
	//定义一个进位数的指针
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		//如果l1 不等于null时，就取他的值，等于null时，就赋值0，保持两个链表具有相同的位数
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		} else {
			n1 = 0
		}
		//如果l2 不等于null时，就取他的值，等于null时，就赋值0，保持两个链表具有相同的位数
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		} else {
			n2 = 0
		}
		//将两个链表的值，进行相加，并加上进位数
		sum := n1 + n2 + carry
		//计算进位数
		carry = sum / 10
		//计算两个数的和，此时排除超过10的请况（大于10，取余数）
		sum = sum % 10
		if head == nil {
			head = &ListNode{Val: sum}
			prev = head
		} else {
			prev.Next = &ListNode{Val: sum}
			prev = prev.Next
		}
	}
	//最后判断是不是还有进位
	if carry > 0 {
		prev.Next = &ListNode{Val: carry}
	}
	return head
}
