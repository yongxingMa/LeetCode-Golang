package main

import (
	"fmt"
	"testing"
)

/**
序号：198
标题：打家劫舍
日期：2022/06/15
类型：
*/
func Test198(t *testing.T) {

	fmt.Println(climbStairs(8))
}

/**
动态规划可以用递归或迭代的方法去实现
*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}
