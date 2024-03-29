package LeetCode

import (
	"fmt"
	"sort"
	"testing"
)

/*
*
序号：015
标题：三数之和
日期：2022/05/31
*/
func Test015(t *testing.T) {
	//var height = []int{3, 0, -2, -1, 1, 2}
	var height = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum2(height))
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	fmt.Println(nums)
	//ans := make([][]int, 0)
	ans := [][]int{}
	//first 是第一个指针，目标值
	//second 是第二个指针，向右移动
	//third 是第三个指针，向左移动
	for first := 0; first < n; first++ {
		fmt.Println(nums[first])
		//避免重复
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		target := -1 * nums[first]
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			//保证b在c的左侧，大于目标值，c往左侧移动，循环
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			//重复后，说明没有答案
			if second == third {
				break
			}
			//等于目标值，记录当前值
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := len(nums) - 1
		target := -1 * nums[first]
		//双指针判断等于target的
		for second := first + 1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			//third往左侧移动
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			// 等于目标值后，记录当前值
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
