package main

import (
	"fmt"
	"testing"
)

/**
序号：009
标题：乘积小于 K 的子数组
日期：2022/12/23
*/

func Test009(t *testing.T) {
	target := 100
	nums := []int{10, 5, 2, 6}
	fmt.Println(numSubarrayProductLessThanK2(nums, target))
}

func numSubarrayProductLessThanK2(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	left, right := 0, 0
	total := 1
	cnt := 0
	for ; right < n; right++ {
		total *= nums[right]

		for total >= k && left <= right {
			total /= nums[left]
			left++
		}
		//count直接计算中间有几种可能
		cnt += right - left + 1
	}
	return cnt
}

//解法一 滑动窗口
func numSubarrayProductLessThanK(nums []int, k int) int {
	start, end := 0, 0
	ans := 0
	sum := 1
	for start <= end && end <= len(nums)-1 {
		sum = sum * nums[end]
		//乘积小于k，需要保存
		if sum < k {
			//增加次数
			ans++
			//滑动窗口右侧右移
			if end < len(nums)-1 {
				end++
			} else {
				sum = 1
				start++
				//复原 从第二个开始
				end = start
			}
		} else {
			sum = 1
			start++
			//复原 从第二个开始
			end = start
		}
	}
	//滑动窗口，一位数保存，两位数保存，
	return ans
}
