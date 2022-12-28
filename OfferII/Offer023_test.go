package main

import (
	"fmt"
	"testing"
)

/**
序号：023
标题：两个链表的第一个重合节点
日期：2022/12/27
*/

func Test023(t *testing.T) {
	s := "abba"
	fmt.Println(countSubstrings(s))
}

//解法一 哈希表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hashmap := map[*ListNode]bool{}
	for headB != nil {
		hashmap[headB] = true
		headB = headB.Next
	}
	for headA != nil {
		if hashmap[headA] == true {
			return headA
		} else {
			headA = headA.Next
		}
	}
	return nil
}

//解法二 双指针
//每步操作需要同时更新指针pA 和 pB。
//如果指针pA 不为空，则将指针 pA 移到下一个节点；如果指针 pB 不为空，则将指针 pB 移到下一个节点。
//如果指针 pA 为空，则将指针 pA 移到链表 headB 的头节点；如果指针 pB 为空，则将指针 pB 移到链表 headA 的头节点。
//当指针 pA 和 pB 指向同一个节点或者都为空时，返回它们指向的节点或者 nil。
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		// 当a链表为空，从b链表开始遍历
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		// 当b链表为空，从a链表开始遍历
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	//如果pa=pb。说明为交点链表，或同时为nil
	return pa
}
