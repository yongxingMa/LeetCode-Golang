package main

import (
	"fmt"
	"testing"
)

/**
序号：35
标题：二分查找
日期：2022/06/02
类型：
*/
func Test35(t *testing.T) {
	var height = []int{3, 0, -2, -1, 1, 2}
	fmt.Println(threeSum(height))
}

/**

 */
func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums) - 1
	for start <= end {
		mid := (end-start)/2 + start
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return end + 1
}
