package main

import (
	"fmt"
	"testing"
)

/**
序号：058
标题：左旋转字符串
日期：2023/01/04
*/

func Test058(t *testing.T) {
	s := "abba"
	fmt.Println(countSubstrings(s))
}

//解法一 数组拼接
func reverseLeftWords(s string, n int) string {
	ans := []byte{}
	for i := n; i < len(s); i++ {
		ans = append(ans, s[i])
	}
	for i := 0; i < n; i++ {
		ans = append(ans, s[i])
	}
	return string(ans)
}

//解法二 不需要额外空间
func reverseLeftWords2(s string, n int) string {
	//把前边的拼接到后边
	for i := 0; i < n; i++ {
		s = s + string(s[i])
	}
	return s[n:]
}
