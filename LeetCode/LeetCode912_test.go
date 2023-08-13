package LeetCode

import (
	"fmt"
	"math/rand"
	"testing"
)

/*
*
序号：912
标题：排序算法
日期：2023/08/12
类型：
*/
func Test912(t *testing.T) {
	var nums = []int{5, 1, 1, 2, 0, 0}
	fmt.Println(sortArray1(nums))
}

// 快速排序
func sortArray1(nums []int) []int {
	QuickSort(nums, 0, len(nums)-1)
	return nums
}

func QuickSort(nums []int, start, end int) {
	if start < end {
		pivotIndex := partition(nums, start, end)
		QuickSort(nums, start, pivotIndex-1)
		QuickSort(nums, pivotIndex+1, end)
	}
}

func partition(nums []int, left, right int) int {
	// 随机选取基准数
	p := rand.Intn(right-left+1) + left
	nums[left], nums[p] = nums[p], nums[left]
	for left < right {
		for nums[left] < nums[right] && left < right {
			// 修改原来的 nums[j] >= nums[i]，增加交换频率
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		for nums[left] < nums[right] && left < right {
			// 修改原来的 nums[j] >= nums[i]，增加交换频率
			left++
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}
	return left
}
