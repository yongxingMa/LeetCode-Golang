package LeetCode

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
这道题目的前提是数组为有序数组，同时题目还强调数组中无重复元素，因为一旦有重复元素，使用二分查找法返回的元素下标可能不是唯一的，这些都是使用二分法的前提条件，
1、左闭右开
在区间[left, right)，使用while (left < right),因为最右边的right取不到
if (nums[middle] > target) {right = middle}，因为left < right，下一个循环取到了nums[middle-1]
2、左闭右闭
在区间[left, right]，使用while (left <= right)，右边right可以取到
if (nums[middle] > target) {right = middle - 1}，因为nums[middle]已经判断过了
*/
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low < high {
		//注：这里要加low
		mid := (high-low)/2 + low
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1
}
