package main

import (
	"fmt"
	"testing"
)

/**
序号：232
标题：用栈实现队列
日期：2022/06/12
类型：
*/
func Test232(t *testing.T) {
	var nums = []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

/**
void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false
*/
type MyQueue struct {
	//输入栈、输出栈
	inStack, outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}

}

//队列添加元素，只从输入栈输入
func (q *MyQueue) Push(x int) {
	//往输入栈添加元素
	q.inStack = append(q.inStack, x)

}

//队列弹出元素
func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	//输出栈的最后一个弹出
	x := q.outStack[len(q.outStack)-1]
	//删掉最后一个元素
	q.outStack = q.outStack[:len(q.outStack)-1]
	return x
}

//放到输出栈
func (q *MyQueue) in2out() {
	for len(q.inStack) > 0 {
		//把输入栈最后一个，放到输出栈
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		//输入栈的最后一个删掉
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0

}
