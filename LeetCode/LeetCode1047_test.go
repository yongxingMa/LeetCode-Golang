package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：1047
标题：删除字符串中的所有相邻重复项
日期：2023/01/09
*/
func Test1047(t *testing.T) {
	var s = "abbaca"
	fmt.Println(removeDuplicates(s))
}

func removeDuplicates(s string) string {
	//定义需要匹配的字符
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		// 	栈不空 且 与栈顶元素相等
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	//最后stack中如果有值，说明有没有配对的
	return string(stack)
}
