package main

import (
	"fmt"
	"testing"
)

/**
序号：977
标题：有序数组的平方
日期：2022/06/03
类型：
*/
func Test167(t *testing.T) {
	var nums = []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

/**
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
*/
func twoSum067(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1
	for low < high {
		sum := numbers[low] + numbers[high]
		if sum == target {
			return []int{low + 1, high + 1}
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return []int{-1, -1}
}
