package LeetCode

import (
	"fmt"
	"math"
	"testing"
)

/**
序号：322
标题：零钱兑换
日期：2022/06/04
类型：
*/
func Test322(t *testing.T) {
	var nums = []int{1, 2, 5}
	fmt.Println(coinChange(nums, 11))
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	//初始化
	dp[0] = 0
	// 初始化为math.MaxInt32
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}
	//遍历物品,i从1开始，从0开始没
	for i := 0; i < len(coins); i++ {
		// 遍历背包
		for j := coins[i]; j <= amount; j++ {
			if j >= coins[i] {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
				fmt.Println(dp, j, i)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化dp[0]
	dp[0] = 0
	// 初始化为math.MaxInt32
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}

	// 遍历物品
	for i := 0; i < len(coins); i++ {
		// 遍历背包
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				// 推导公式
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
				fmt.Println(dp, j, i)
			}
		}
	}
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
