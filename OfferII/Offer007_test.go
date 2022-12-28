package main

import (
	"fmt"
	"sort"
	"testing"
)

/**
序号：007
标题：数组中和为 0 的三个数
日期：2022/12/23
*/

func Test007(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{0, 0, 0}
	fmt.Println(threeSum3(nums))
}

//解法一 暴力法 失败
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	hashmap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			ano := -(nums[i] + nums[j])
			if p, ok := hashmap[ano]; ok {
				if p >= j {
					ans = append(ans, []int{nums[i], nums[j], nums[p]})
				}
				break
			} else {
				hashmap[nums[j]] = j
				hashmap[nums[i]] = i
			}
		}
	}
	return ans
}

//解法二 双指针
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for a := 0; a < len(nums); a++ {
		//如果重复的直接跳过
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		//第三个值，从尾部开始
		c := len(nums) - 1
		//目标值 取反数
		target := -1 * nums[a]
		for b := a + 1; b < len(nums); b++ {
			//跳过重复的值
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			//保证b在c的左侧，大于目标值，c往左侧移动
			for b < c && nums[b]+nums[c] > target {
				c--
			}
			//重复后，说明没有答案
			if b == c {
				break
			}
			//等于目标值，记录当前值
			if nums[b]+nums[c] == target {
				ans = append(ans, []int{nums[a], nums[b], nums[c]})
			}
		}
	}
	return ans
}

//解法二 双指针
func threeSum3(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	//a b c 目标值 左指针 右指针
	for a := 0; a < len(nums); a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		c := len(nums) - 1
		target := -1 * nums[a]
		for b := a + 1; b < len(nums); b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			for b < c && nums[b]+nums[c] > target {
				c--
			}
			if b == c {
				break
			}
			if nums[b]+nums[c] == target {
				ans = append(ans, []int{nums[a], nums[b], nums[c]})
			}
		}
	}
	return ans
}
