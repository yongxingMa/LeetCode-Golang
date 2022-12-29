package LeetCode

import (
	"fmt"
	"math"
	"testing"
)

/**
序号：034
标题：搜索插入位置
日期：2022/06/02
*/
func Test034(t *testing.T) {
	var height = []int{7, 8, 8, 8, 9, 10}
	fmt.Println(searchRange3(height, 8))
}

// 暴力解法
func searchRange(nums []int, target int) []int {
	ans := make([]int, len(nums))
	first := math.MaxInt32
	last := math.MinInt32
	//判断是否存在target，存在直接返回下标
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			first = min(first, i)
			last = max(last, i)
		}
	}
	if first == math.MaxInt32 && last == math.MinInt32 {
		ans = []int{-1, -1}
	} else {
		ans = []int{first, last}
	}
	return ans
}

// 双指针
func searchRange2(nums []int, target int) []int {
	ans := []int{-1, -1}
	// 左右指针
	left, right := 0, len(nums)-1
	//判断是否存在target，存在直接返回下标
	for left <= right {
		if nums[right] != target {
			right--
		} else {
			if nums[left] != target {
				left++
			} else {
				break
			}
		}
	}
	if right != -1 {
		ans = []int{left, right}
	}
	return ans
}

// 二分查找
// 找到第一个大于等于target的下标，设为start，那么start下标对应的值就应该是第一次出现target的下标。
// 找到第一个大于等于target+1的下标，设为end，那么end-1下标就是target最后一次出现的下标，最终返回[start, end - 1]。
func searchRange3(nums []int, target int) []int {
	//二分查找,找到第一个目标target的值
	start := search2(nums, target)
	//if start == len(nums) || nums[start] != target {
	//	return []int{-1, -1}
	//}
	// 如果 start 存在，那么 end 必定存在
	end := search2(nums, target+1)
	return []int{start, end - 1}
}

//查找第一个大于等于target的下标
func search2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		// 这里改成 nums[mid] > target 表示查找第一个小于或等于target的下标
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
