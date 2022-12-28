package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：167
标题：两数之和 II - 输入有序数组
日期：2022/06/03
类型：
*/
func Test167(t *testing.T) {
	var nums = []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

//双指针解法
func twoSum067(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1
	for low < high {
		sum := numbers[low] + numbers[high]
		if sum == target {
			//数组中的第几位，不是数组下标，不是从0开始的，从1开始
			return []int{low + 1, high + 1}
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return nil
}
