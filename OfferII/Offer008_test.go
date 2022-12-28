package main

import (
	"fmt"
	"math"
	"testing"
)

/**
序号：008
标题：和大于等于 target 的最短子数组
日期：2022/12/23
*/

func Test008(t *testing.T) {
	target := 11
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(minSubArrayLen(target, nums))
}

//解法一 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	sum := 0
	ans := math.MaxInt32
	start, end := 0, 0
	for end < len(nums) {
		sum = sum + nums[end]
		for sum >= target {
			//记录当前的答案，start到end的距离
			ans = min(ans, end-start+1)
			//窗口最左侧数据删掉
			sum = sum - nums[start]
			//窗口左侧右移
			start++
		}
		//窗口右侧右移
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
