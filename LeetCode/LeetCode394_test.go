package LeetCode

import (
	"fmt"
	"strings"
	"testing"
)

/**
序号：394
标题：字符串解码
日期：2022/06/12
类型：
*/

/*
*
void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false
*/
func Test394(t *testing.T) {
	ransomNote := "100[leetcode]"
	fmt.Println(decodeString(ransomNote))
}

// 递归解法
func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	num := 0
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] < '9' {
			// 将字符串转换为整数类型
			//n, _ := strconv.Atoi(string(s[i]))
			n := int(s[i] - '0')
			num = num*10 + n
		} else if s[i] == '[' {
			strStack = append(strStack, result)
			result = ""
			numStack = append(numStack, num)
			num = 0
		} else if s[i] == ']' {
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			result = str + strings.Repeat(result, count)
		} else {
			result += string(s[i])
		}
	}
	return result
}
