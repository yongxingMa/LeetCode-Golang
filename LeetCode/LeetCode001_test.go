package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：001
标题：两数之和
日期：2022/06/11
类型：/简单
*/
func Test001(t *testing.T) {
	var nums = []int{3, 3}
	fmt.Println(twoSum001(nums, 6))
}

func twoSum001(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		another := target - nums[i]
		//ok 是存在
		if p, ok := hashTable[another]; ok {
			return []int{p, i}
		} else {
			//key是数据，value是第几个值
			hashTable[x] = i
		}

	}
	return nil
}

func twoSum002(nums []int, target int) []int {
	hashmap := map[int]int{}
	for i, num := range nums {
		ano := target - num
		if p, ok := hashmap[ano]; ok {
			// 返回存在的下标值和当前的下标指
			return []int{p, i}
		} else {
			hashmap[num] = i
		}
	}
	return []int{}
}
