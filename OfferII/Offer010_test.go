package main

import (
	"fmt"
	"testing"
)

/**
序号：010
标题：和为 k 的子数组
日期：2022/12/24
*/

func Test010(t *testing.T) {
	target := 7
	nums := []int{3, 4, 7, 2, -3, 1, 4, 2}
	fmt.Println(subarraySum(nums, target))
}

//前缀和 + 哈希表
func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	hashmap := map[int]int{}
	// 这里是 sum == k 的情况下 sum - k = 0，计数器要 + 1，所以默认为1
	hashmap[0] = 1
	for i := 0; i < len(nums); i++ {
		//存储前缀和
		sum += nums[i]
		// 关键： 当前的前缀和与目标值进行比较，
		// 如果当前的前缀和与目标值相等，说明从0开始到i的和满足要求
		// 如果当前的前缀和比目标值大，寻找hashmap中是否存在sum-k的值，如果存在，说明有sum-k这个坐标到i的和是满足要求的
		// 如果当前的前缀和比目标值小，找不到数据
		if _, ok := hashmap[sum-k]; ok {
			// 计数器累加
			count += hashmap[sum-k]
		}
		// 前缀和可能会重复，存在两个相同的前缀和，这是两个答案，都要累加
		hashmap[sum] += 1
	}
	return count
}
