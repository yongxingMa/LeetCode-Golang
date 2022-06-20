package main

import (
	"fmt"
	"testing"
)

/**
序号：704
标题：二分查找
日期：2022/06/02
类型：
*/
func Test704(t *testing.T) {
	var height = []int{3, 0, -2, -1, 1, 2}
	fmt.Println(threeSum(height))
}

/**
先从中间值开始比较，
*/
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		//mid中间值坐标
		mid := (high-low)/2 + low
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
