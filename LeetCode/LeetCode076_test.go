package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：76
标题：最小覆盖子串
日期：2023/08/29
类型：
*/
func Test076(t *testing.T) {
	fmt.Println(climbStairs(8))
}

// 双指针滑动窗口
func minWindow(s string, t string) string {
	sLen, tLen := len(s), len(t)
	// 先遍历 t， 记录元素到哈希表 need 中；记录需要遍历的个数 len(t) 到 needCnt。
	need := make(map[byte]int, tLen)
	for i := 0; i < tLen; i++ {
		need[t[i]]++
	}
	res := ""
	l, min := 0, sLen+1
	match := 0
	window := make(map[byte]int)
	for r := 0; r < sLen; r++ {
		// 1. 直接将s[right]加入到区间，形成（left, right]
		b := s[r]
		window[b]++
		if window[b] == need[b] {
			match++
		}
		for match == len(need) {
			if r-l < min {
				min = r - l + 1
				res = s[l : r+1]
			}
			b = s[l]
			if window[b] == need[b] {
				match--
			}
			window[b]--
			l++
		}
	}
	return res
}
