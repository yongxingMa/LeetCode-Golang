package main

import (
	"fmt"
	"testing"
)

/**
序号：33
标题：搜索旋转排序数组
日期：2022/06/02
类型：二分查找
*/
func Test33(t *testing.T) {
	var height = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search33(height, 3))
}

/**

 */
func search33(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		//mid中间值坐标
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[low] { // 中间值比最左边的大，说明左边是有序的
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] { // 中间值比最右边的小，说明右边是有序的
			if nums[mid] <= target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return -1
}
