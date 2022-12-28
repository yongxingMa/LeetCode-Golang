package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：136
标题：只出现一次的数字
日期：2022/06/19
类型：滑动窗口/中等
*/
func Test136(t *testing.T) {
	var nums = []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	hashmap := make(map[int]int)
	for _, num := range nums {
		if hashmap[num] != 0 {
			hashmap[num] = 0
		} else {
			hashmap[num] = num
		}
	}
	for i, num := range hashmap {
		if hashmap[i] != 0 {
			return num
		}
	}
	return 0
}

//异或算法
//func singleNumber(nums []int) int {
//	single := 0
//	for _, num := range nums {
////数组中的全部元素的异或运算结果即为数组中只出现一次的数字。
//		single ^= num
//	}
//	return single
//}
