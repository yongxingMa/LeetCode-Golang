package main

import (
	"fmt"
	"testing"
)

/**
序号：283
标题：移动零
日期：2022/06/01
双指针解法
*/
func Test283(t *testing.T) {
	var height = []int{3, 0, -2, -1, 1, 2}
	fmt.Println(threeSum(height))
}

/**
：定义left、right指针，right从左往右移动，遇上非0元素，交换left和right对应的元素，交换之后left++
0 1 0 3 12
1 0 0 3 12
1 3 0 0 12
1 3 12 0 0
*/
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
