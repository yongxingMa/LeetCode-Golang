package main

import (
	"fmt"
	"testing"
)

/**
序号：977
标题：有序数组的平方
日期：2022/06/03
类型：
*/
func Test977(t *testing.T) {
	var nums = []int{-4,-1,0,3,10}
	fmt.Println(sortedSquares(nums))
}

/**
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
 */
func sortedSquares(nums []int) []int {
	//解法一
	//ans:= make([]int, len(nums))
	//for i, v := range nums {
	//	ans[i] = v * v
	//}
	//sort.Ints(ans)
	//return ans

	//解法二
	n := len(nums)
	lastNegIndex := -1
	//找到最后一个负数
	for i := 0; i < n && nums[i] < 0; i++ {
		lastNegIndex = i
	}
	//从0到n设置一个数组
	ans := make([]int, 0, n)
	i := lastNegIndex
	j := lastNegIndex + 1
	for i >= 0 || j < n {
		if i < 0 { //说明没有负数，都是正数
			ans = append(ans, nums[j]*nums[j])
			j++
		} else if j == n { //如果最后一个是负数
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