package main

import (
	"fmt"
	"testing"
)

/**
序号：014
标题：字符串中的变位词
日期：2022/12/25
*/

func Test014(t *testing.T) {
	s1 := "ab"
	//s2 := "eidbaooo"
	s2 := "bcaba"
	fmt.Println(checkInclusion(s1, s2))
}

//滑动窗口+哈希表
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	// 数组模拟哈希表
	var cnt1, cnt2 [26]int
	for i, _ := range s1 {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		// 上边for循环已经比较过来前边s1的长度的值，这里从s1的长度右边一个开始比较
		// s2[i]表示s1字符串长度的右边一个字母，滑动窗口右侧+1
		cnt2[s2[i]-'a']++
		// s2[i-len(s1)] 表示s1字符串的第一个字母 滑动窗口左侧-1
		cnt2[s2[i-len(s1)]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}
