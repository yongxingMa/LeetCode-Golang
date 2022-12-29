package LeetCode

import (
	"fmt"
	"sort"
	"testing"
)

/**
序号：977
标题：有序数组的平方
日期：2022/06/03
类型：
*/
func Test977(t *testing.T) {
	var nums = []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares3(nums))
}

//解法一 暴力解法
func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	for i, v := range nums {
		ans[i] = v * v
	}
	sort.Ints(ans)
	return ans
}

//解法二 双指针
// nums = [-4,-1,0,3,10]
func sortedSquares2(nums []int) []int {

	n := len(nums)
	lastNegIndex := -1
	//找到最后一个负数
	for i := 0; i < n && nums[i] < 0; i++ {
		lastNegIndex = i
	}
	//make([]int, 0, n) 初始化 len=0 cap=n
	//make([]int, n) 初始化 len=n cap=n
	ans := make([]int, 0, n)
	i := lastNegIndex
	j := lastNegIndex + 1
	for i >= 0 || j < n {
		if i < 0 { //说明没有负数，都是正数
			ans = append(ans, nums[j]*nums[j])
			j++
		} else if j == n { //说明i是数组最后一个数
			ans = append(ans, nums[i]*nums[i])
			i--
		} else if nums[i]*nums[i] < nums[j]*nums[j] {
			ans = append(ans, nums[i]*nums[i])
			i--
		} else {
			ans = append(ans, nums[j]*nums[j])
			j++
		}
	}
	return ans
}

//解法三 双指针优化 不需要判断临界点，从左右两边最大值判断，加入数组的最后一位
func sortedSquares3(nums []int) []int {
	n := len(nums)
	//从0到n初始化一个数组
	ans := make([]int, n)
	i, j := 0, n-1
	for pos := n - 1; pos >= 0; pos-- {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			ans[pos] = nums[i] * nums[i]
			i++
		} else {
			ans[pos] = nums[j] * nums[j]
			j--
		}
	}
	return ans
}
