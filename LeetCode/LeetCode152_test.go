package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：152
标题：乘积最大子数组
日期：2023/07/30
*/
func Test152(t *testing.T) {
	nums := []int{-2, 3, -4}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}
