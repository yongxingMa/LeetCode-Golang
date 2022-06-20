package main

import (
	"fmt"
	"testing"
)

/**
序号：020
标题：有效的括号
日期：2022/06/19
类型：滑动窗口/中等
*/
func Test020(t *testing.T) {
	var s = "()[]{}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	//定义需要匹配的字符
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		//如果存在右边的括号，说明也有左边的括号，找不到直接返回false，最后stack中如果有值，说明有没有配对的
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			//弹出栈
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
