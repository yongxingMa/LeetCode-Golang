package main

import (
	"fmt"
	"testing"
)

/**
序号：011
标题：0 和 1 个数相同的子数组
日期：2022/12/24
*/

func Test011(t *testing.T) {
	nums := []int{0, 1, 0, 1, 3, -3}
	fmt.Println(findMaxLength(nums))
}

//前缀和 + 哈希表
//题干中前提条件 nums[i] 不是 0 就是 1
func findMaxLength(nums []int) int {
	count, sum := 0, 0
	hashmap := map[int]int{}
	// 这里是 sum == k 的情况下 sum - k = 0，计数器要 + 1，所以默认为1
	hashmap[0] = -1
	for i := 0; i < len(nums); i++ {
		// 把数字0替换为-1
		if nums[i] == 0 {
			nums[i] = -1
		}
		//存储前缀和
		sum += nums[i]
		//关键： 如果出现了相同的前缀和，说明从现在的i到-->hashmap存储的相同的前缀和的i的这个长度为0，长度为 i-hashmap[sum]
		// 题干中有个前提条件 nums[i] 不是 0 就是 1
		if _, ok := hashmap[sum]; ok {
			count = max(count, i-hashmap[sum])
		} else {
			// 只保存第一次出现的前缀和对应的i的坐标
			hashmap[sum] = i
		}
	}
	return count
}
