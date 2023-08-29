package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：560
标题：和为K的子数组
日期：2023/08/28
类型：子串
*/

func Test560(t *testing.T) {
	nums := []int{1, 2, 3, 6}
	fmt.Println(subarraySum(nums, 6))
}

// 为什么这题不可以用双指针/滑动窗口：因为nums[i]可以小于0，也就是说右指针i向后移1位不能保证区间会增大，左指针j向后移1位也不能保证区间和会减小。给定j，i的位置没有二段性。
// 前缀和 + 哈希表
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	// 初始化{0，1}是为了防止从序号0开始加到i时刚好为k的情况出现时，没法找到对应数值计数
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}

	return count
}
