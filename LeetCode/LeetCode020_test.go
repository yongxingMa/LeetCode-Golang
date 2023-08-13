package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：020
标题：有效的括号
日期：2022/06/19
类型：堆栈
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
	// 遇到左括号
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		//大于0表示右括号，需要在栈里找左括号，找不到直接返回false，
		if pairs[s[i]] > 0 {
			//这里必须在栈顶上找到对应的左括号，不然也是不合法的
			//stack[len(stack)-1] 表示最后一个
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			//在栈顶找到后，弹出左侧括号的栈
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	//最后stack中如果有值，说明有没有配对的
	return len(stack) == 0
}

func isValid2(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	hashmap := map[byte]byte{'}': '{', ']': '[', ')': '('}
	stack := []byte{}
	for i := 0; i < n; i++ {
		// 右括号，需要在栈里找左括号，找不到直接返回false，
		if s[i] == ']' || s[i] == '}' || s[i] == ')' {
			//这里必须在栈顶上找到对应的左括号，不然也是不合法的
			if len(stack) == 0 || stack[len(stack)-1] != hashmap[s[i]] {
				return false
			}
			//在栈顶找到后，弹出左侧括号的栈
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
