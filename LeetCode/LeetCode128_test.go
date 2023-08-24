package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：128
标题：最长连续序列
日期：2023/08/24
类型：
*/
func Test128(t *testing.T) {
	nums := []int{1, 2, 0, 1}
	fmt.Println(longestConsecutive1(nums))
}

// 哈希表 先去重
func longestConsecutive1(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	//记录一个最大值
	maxNum := 0
	for num := range numSet {
		// 因为是无序的，这里要判断他的上一个值是否存在，存在就回滚，不存在继续
		if !numSet[num-1] {
			currentNum := num
			// 记录长度的值
			currentStreak := 1
			// 当前值为true，寻找下一个是否为true，为true表示存在
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if currentStreak > maxNum {
				maxNum = currentStreak
			}
		}
	}
	return maxNum
}
