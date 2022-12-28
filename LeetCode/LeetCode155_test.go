package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：115
标题：最小栈
日期：2022/06/10
类型：滑动窗口/中等
*/
func Test最小栈(t *testing.T) {
	var s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

//需要一个辅助栈，用来存储最小值
type MinStack struct {
	stack    []int
	minStack []int
}

//func Constructor() MinStack {
//	return MinStack{
//		stack:    []int{},
//		minStack: []int{math.MaxInt64},
//	}
//}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(val, top))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]

}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
