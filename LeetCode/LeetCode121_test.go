package LeetCode

import (
	"fmt"
	"math"
	"testing"
)

/*
*
序号：121
标题：买卖股票的最佳时机
日期：2023/07/31
类型：
*/
func Test121(t *testing.T) {
	fmt.Println(climbStairs(8))
}

/*
*
动态规划可以用递归或迭代的方法去实现
*/
//暴力循环
func maxProfit(prices []int) int {
	// 记录最大利润
	maxprofit := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			//第i天买，遍历后边每天一次都卖，看哪天最赚钱
			profit := prices[j] - prices[i]
			if profit > maxprofit {
				maxprofit = profit
			}
		}
	}
	return maxprofit
}

// 一次遍历
// 记录一个最小买进的值，每次遍历记录最大值
func maxProfit2(prices []int) int {
	minprice := math.MaxInt
	maxprofit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i]-minprice > maxprofit {
			maxprofit = prices[i] - minprice
		}
	}
	return maxprofit
}
