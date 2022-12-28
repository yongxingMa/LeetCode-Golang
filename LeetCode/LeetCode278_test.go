package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：278
标题：第一个错误的版本
日期：2022/06/02
双指针解法
*/
func Test278(t *testing.T) {
	var height = []int{3, 0, -2, -1, 1, 2}
	fmt.Println(threeSum(height))
}

/**
：定义left、right指针，right从左往右移动，遇上非0元素，交换left和right对应的元素，交换之后left++
*/
func firstBadVersion(n int) int {
	return binarySearch(1, n)
}

func binarySearch(start int, end int) int {
	if start == end {
		return start
	}
	mid := (end-start)/2 + start

	if isBadVersion(mid) == true {
		return binarySearch(start, mid)
	} else {
		return binarySearch(mid+1, end)
	}
}

func isBadVersion(n int) bool {
	return true
}
