package main

import (
	"fmt"
	"testing"
)

/**
序号：012
标题：左右两边子数组的和相等
日期：2022/12/24
*/

func Test012(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
}

//
func pivotIndex(nums []int) int {
	total := 0
	sum := 0
	for _, num := range nums {
		// 记录总和
		total += num
	}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// 全部之和 减去 0-i的和 是右边的和
		// 0-i的和减去i 是左边的和
		if total-sum == sum-nums[i] {
			return i
		}
	}
	return -1
}
