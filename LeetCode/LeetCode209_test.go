package LeetCode

import (
	"fmt"
	"math"
	"testing"
)

/**
序号：209
标题：长度最小的子数组
日期：2022/12/29
*/
func Test209(t *testing.T) {
	var height = []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen1(7, height))
}

//解法一 暴力解法
func minSubArrayLen1(target int, nums []int) int {
	ans := math.MaxInt32

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				ans = min(ans, j-i+1)
				break
			}
		}
	}
	return ans
}

//解法二 双指针
func minSubArrayLen2(target int, nums []int) int {
	left, right := 0, 0
	sum := 0
	ans := math.MaxInt32
	// 这里要注意，两个for循环，第一层移动右窗口，第二层移动左窗口
	// 窗口里的值满足sum >= target 就要循环左窗口（这里一定不要用if来判断）
	// 窗口里的值满足sum < target,就循环右窗口
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
