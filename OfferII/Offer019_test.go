package main

import (
	"fmt"
	"testing"
)

/**
序号：019
标题：最多删除一个字符得到回文
日期：2022/12/27
*/

func Test019(t *testing.T) {
	s := "abc"
	fmt.Println(validPalindrome(s))
}

// 双指针+递归解法
func validPalindrome(s string) bool {
	flag := false
	var check func(left, right int) bool

	check = func(left, right int) bool {
		for left < right {
			if s[left] == s[right] {
				left++
				right--
			} else if flag == true {
				return false
			} else {
				//一次删除的机会判断
				flag = true
				return check(left+1, right) || check(left, right-1)
			}
		}
		return true
	}
	// 全部字符串，调用上边的方法
	return check(0, len(s)-1)
}
