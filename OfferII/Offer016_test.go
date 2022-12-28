package main

import (
	"fmt"
	"testing"
)

/**
序号：016
标题：不含重复字符的最长子字符串
日期：2022/12/25
*/

func Test016(t *testing.T) {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

//滑动窗口
func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	hashmap := map[byte]int{}
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	right, ans := -1, 0
	if len(s) == 0 {
		return 0
	}
	// 循环i，也就是左指针
	for i := 0; i < len(s); i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(hashmap, s[i-1])
		}
		//判断是否存在重复的字符已经存在，如果==0 说明不存在 进行+1操作
		for right+1 < len(s) && hashmap[s[right+1]] == 0 {
			hashmap[s[right+1]]++
			// 右指针向右移动
			right++
		}
		// 右指针减去左指针的长度
		ans = max(ans, right-i+1)
	}
	return ans
}
