package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：707
标题：设计链表
日期：2022/12/30
类型：
*/
func Test707(t *testing.T) {
	var height = []int{3, 0, -2, -1, 1, 2}
	fmt.Println(threeSum(height))
}

type MyLinkedList struct {
	//虚拟头节点
	head *ListNode
	//容量
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{&ListNode{}, 0}
}

//下标从0开始的，head是空节点 val=0
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	pred := this.head
	for i := 0; i <= index; i++ {
		pred = pred.Next
	}
	return pred.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// 先判断是不是超出长度
	if index > this.size {
		return
	}
	//如果index是负数的情况
	if index < 0 {
		index = 0
	}
	//长度+1
	this.size++
	pred := this.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	toAdd := &ListNode{val, pred.Next}
	pred.Next = toAdd
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	// 先判断是不是超出长度
	if index < 0 || index >= this.size {
		return
	}
	//长度-1
	this.size--
	pred := this.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	pred.Next = pred.Next.Next
}
