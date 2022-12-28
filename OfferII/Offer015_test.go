package main

import (
	"fmt"
	"testing"
)

/**
序号：015
标题：字符串中的所有变位词
日期：2022/12/25
*/

func Test015(t *testing.T) {
	s := "aa"
	p := "bb"
	fmt.Println(findAnagrams(s, p))
}

//滑动窗口
func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	if len(p) > len(s) {
		return ans
	}
	// 数组模拟哈希表
	var cnt1, cnt2 [26]int
	for i, _ := range p {
		// 短的单词
		cnt1[p[i]-'a']++
		// 长的单词
		cnt2[s[i]-'a']++
	}
	if cnt1 == cnt2 {
		ans = append(ans, 0)
	}
	for i := len(p); i < len(s); i++ {
		// 上边for循环已经比较过来前边p的长度的值，这里从p的长度右边一个开始比较
		// s[i]表示p字符串长度的右边一个字母，滑动窗口右侧+1
		cnt2[s[i]-'a']++
		// s2[i-len(p)] 表示p字符串的第一个字母 滑动窗口左侧-1
		cnt2[s[i-len(p)]-'a']--
		if cnt1 == cnt2 {
			// 滑动窗口左侧 + 1才是启始值
			ans = append(ans, i-len(p)+1)
		}
	}
	return ans
}
