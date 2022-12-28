package LeetCode

import (
	"testing"
)

/**
序号：344
标题：反转字符串
日期：2022/06/05
类型：双指针
难度：简单
*/
func Test344(t *testing.T) {
	//var nums = []byte{"h","e","l","l","o"}
	//fmt.Println(sortedSquares(nums))
}

/**
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
*/
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
