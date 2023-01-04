package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：541
标题：反转字符串 II
日期：2023/0104
*/
func Test541(t *testing.T) {
	s := "abcdefg"
	k := 2
	fmt.Println(reverseStr(s, k))
}

// 解法一 模拟
// 反转每个下标从 2k 的倍数开始的，长度为 k 的子串。若该子串长度不足 k，则反转整个子串。
func reverseStr(s string, k int) string {
	t := []byte(s)
	// 遍历字符串
	for i := 0; i < len(s); i = i + 2*k {
		// min(i+k, len(s)）判断剩余字符串
		sub := t[i:min(i+k, len(s))]
		for j, n := 0, len(sub); j < n/2; j++ {
			sub[j], sub[n-1-j] = sub[n-1-j], sub[j]
		}
	}
	return string(t)
}

//解法二
func reverseStr2(s string, k int) string {
	ss := []byte(s)
	for i := 0; i < len(s); i = i + 2*k {
		if i+k <= len(s) {
			reverse4(ss[i : i+k])
		} else {
			reverse4(ss[i:len(s)])
		}
	}
	return string(ss)
}

//反转字符串
func reverse4(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
