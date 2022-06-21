package main

import (
	"fmt"
	"testing"
)

/**
序号：005
标题：最长回文子串
日期：2022/06/20
类型：经典题目，就是有点难
*/
func Test005(*testing.T) {
	var nums = []int{3, 3}
	fmt.Println(twoSum001(nums, 6))
}

//动态规划解法
func longestPalindrome(s string) string {
	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
