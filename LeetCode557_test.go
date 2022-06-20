package main

import (
	"fmt"
	"testing"
)

/**
序号：557
标题：反转字符串中的单词 III
日期：2022/06/05
类型：双指针
难度：简单
*/
func Test557(t *testing.T) {
	var s = "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}

/**
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
*/

func reverseWords(s string) string {
	length := len(s)
	ret := []byte{}
	for i := 0; i < length; {
		start := i
		//不是空，说明还是一个单词
		for i < length && s[i] != ' ' {
			i++
		}
		for p := start; p < i; p++ {
			ret = append(ret, s[start+i-1-p])
		}
		//说明已经是下一个单词了
		for i < length && s[i] == ' ' {
			i++
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}
