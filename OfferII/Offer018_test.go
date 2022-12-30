package main

import (
	"fmt"
	"strings"
	"testing"
)

/**
序号：018
标题：有效的回文
日期：2022/12/27
*/

func Test018(t *testing.T) {
	s := "123"
	fmt.Println(isPalindrome1(s))
}

//双指针解法
func isPalindrome1(s string) bool {
	//全部转换为小写
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	//这个题目需要过滤其他字符，19题不需要过滤
	for left < right {
		if !judge(s[left]) {
			left++
			continue
		} else if !judge(s[right]) {
			right--
			continue
		} else if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func judge(b byte) bool {
	//如果字节在A-Z或者在0-9之间返回true
	if b >= 'a' && b <= 'z' || b >= '0' && b <= '9' {
		return true
	}
	return false
}
