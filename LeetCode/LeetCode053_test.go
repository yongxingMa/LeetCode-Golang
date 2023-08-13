package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：053
标题：最大子数组和
日期：2022/12/26
请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
nums = [-2,1,-3,4,-1,2,1,-5,4]
*/
func Test053(t *testing.T) {
	nums := []int{1}
	fmt.Println(maxSubArray(nums))
}

// 动态规划
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	//dp := []int{}
	dp[0] = nums[0]
	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		ans = max(ans, dp[i])
	}
	return ans
}
