package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：35
标题：搜索插入位置
日期：2022/06/02
*/
func Test35(t *testing.T) {
	var height = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(height, 0))
}

//暴力解法
func searchInsert(nums []int, target int) int {
	//判断是否存在target，存在直接返回下标
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	// 在头部插入
	if target < nums[0] {
		return 0
	}
	// 在尾部插入
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	// 在中间插入
	orgNums := nums
	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			nums = append(nums[:i-1], target)
			nums = append(orgNums[i:])
			return i
		}
	}
	return -1
}

//双指针解法
func searchInsert2(nums []int, target int) int {
	start, end := 0, len(nums)-1
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
	//如果需要在头部插入，这里end会一直左移，直到end=-1
	//如果需要在尾部插入，start会一直右移，直到start=end
	return end + 1
}
