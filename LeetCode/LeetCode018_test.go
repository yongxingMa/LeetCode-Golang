package LeetCode

import (
	"fmt"
	"sort"
	"testing"
)

/**
序号：018
标题：四数之和
日期：2022/12/23
*/
func Test018(t *testing.T) {
	//var height = []int{1, 0, -1, 0, -2, 2}
	var height = []int{2, 2, 2, 2, 2}
	fmt.Println(fourSum(height, 8))
}

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	//first 是第一个指针，目标值
	//second 是第二个指针，向右移动
	//third 是第三个指针，向左移动
	//zero 新增指针
	for zero := 0; zero < n; zero++ {
		if zero > 0 && nums[zero] == nums[zero-1] {
			continue
		}
		for first := zero + 1; first < n; first++ {
			if first > zero+1 && nums[first] == nums[first-1] {
				continue
			}
			third := n - 1
			targetNew := target - (nums[first] + nums[zero])
			for second := first + 1; second < n; second++ {
				if second > first+1 && nums[second] == nums[second-1] {
					continue
				}
				//保证b在c的左侧，大于目标值，c往左侧移动，循环
				for second < third && nums[second]+nums[third] > targetNew {
					third--
				}
				//重复后，说明没有答案
				if second == third {
					break
				}
				//等于目标值，记录当前值
				if nums[second]+nums[third] == targetNew {
					ans = append(ans, []int{nums[zero], nums[first], nums[second], nums[third]})
				}
			}
		}
	}
	return ans
}
