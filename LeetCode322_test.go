package main

import (
	"fmt"
	"sort"
	"testing"
)

/**
序号：322
标题：零钱兑换
日期：2022/06/04
类型：
*/
func Test322(t *testing.T) {
	var nums = []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

/**
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
*/
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Ints(coins)
	for i := 0; i < len(coins)-1; i++ {
		//amount == c
	}
	return 0
}
