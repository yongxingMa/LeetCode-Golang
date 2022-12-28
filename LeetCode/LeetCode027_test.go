package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：027
标题：移除元素
日期：2022/06/11
*/
func Test027(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 3))
}

//type ListNode struct {
//    Val int
//	Next *ListNode
//}

//暴力解法
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

//双指针的解法
func removeElement2(nums []int, val int) int {
	// right起始在数组最大值右边
	left, right := 0, len(nums)-1
	// left == right有意义
	for left <= right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
