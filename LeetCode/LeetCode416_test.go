package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：416
标题：分割等和子集
日期：2022/12/26
*/
func Test416(t *testing.T) {
	var nums = []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}

func canPartition(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}
	if total%2 == 1 {
		return false
	}
	target := total / 2
	// 创建二维数组 行：物品索引，列：容量（包括 0）
	dp := make([][]bool, len(nums))
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	//循环背包
	for i := 0; i < len(nums); i++ {
		// 第一列全部初始化为true
		dp[i][0] = true
	}
	// 先填表格第 1 行，第 1 个数只能让容积为它自己的背包恰好装满
	if nums[0] <= target {
		dp[0][nums[0]] = true
	}
	for i := 1; i < len(nums); i++ {
		for j := 1; j <= target; j++ {
			if j >= nums[i] {
				// 取一个为true的值
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][target]
}
