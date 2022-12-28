package LeetCode

import (
	"fmt"
	"testing"
)

/**
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
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		//大于0表示，存在右边的括号，说明栈里边也有左边的括号，找不到直接返回false，
		if pairs[s[i]] > 0 {
			//这里必须在栈顶上找到对应的左括号，不然也是不合法的
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
