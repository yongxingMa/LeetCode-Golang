package main

import (
	"fmt"
	"testing"
)

/**
序号：020
标题：回文子字符串的个数
日期：2022/12/27
*/

func Test020(t *testing.T) {
	s := "abba"
	fmt.Println(countSubstrings(s))
}

// 中心扩散
func countSubstrings(s string) int {
	n := len(s)
	cnt := 0
	for i := 0; i < n; i++ {
		// 以i为中心，向两边扩散
		cnt += palindrome(s, i, i)
		// 以i 和 i+1为中心，向两边扩散
		cnt += palindrome(s, i, i+1)
	}
	return cnt
}

func palindrome(s string, start int, end int) int {
	cnt := 0
	for start >= 0 && end < len(s) && s[start] == s[end] {
		cnt++
		start--
		end++
	}
	return cnt
}
