package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：189
标题：轮转数组
日期：2022/06/03
类型：
*/
func Test189(t *testing.T) {
	var nums = []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

/**
给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
*/
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
