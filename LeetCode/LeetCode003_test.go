package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：003
标题：无重复字符的最长子串
日期：2022/06/20
类型：滑动窗口/中等
*/
func Test003(t *testing.T) {
	var s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	hashmap := map[byte]int{}
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	right, ans := -1, 0
	for i := 0; i < len(s); i++ {
		// 当i>0,左指针开始右移
		if i != 0 {
			// 左指针向右移动一格，移除前边的一个字符
			delete(hashmap, s[i-1])
		}
		// 判断右边的数字是不是满足条件，是的话继续右移
		for right+1 < len(s) && hashmap[s[right+1]] == 0 {
			// 记录加入的值
			hashmap[s[right+1]] = 1
			right++
		}
		ans = max(ans, right-i+1)
	}
	return ans
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
