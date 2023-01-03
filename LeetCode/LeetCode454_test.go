package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：454
标题：四数相加II
日期：2023/01/03
*/
func Test454(t *testing.T) {
	var height = []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(height))
}

//解法一：分组 + 哈希表
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	hashmap := map[int]int{}
	ans := 0
	// 遍历nums1和nums2数组，统计两个数组元素之和，和出现的次数，放到hashmap中
	for _, i := range nums1 {
		for _, j := range nums2 {
			hashmap[i+j]++
		}
	}
	// 遍历nums3和nums4数组，找到如果 0-(c+d) 在map中出现过的话，就把map中key对应的value也就是出现次数统计出来
	for _, i := range nums3 {
		for _, j := range nums4 {
			if _, ok := hashmap[0-(i+j)]; ok {
				ans += hashmap[0-(i+j)]
			}
		}
	}
	return ans
}
