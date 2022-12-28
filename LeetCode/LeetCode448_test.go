package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：448
标题：找到所有数组中消失的数字
日期：2022/06/12
双指针解法
*/
func Test448(t *testing.T) {
	var height = []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(height))
}

//【笔记】将所有正数作为数组下标，置对应数组值为负值。那么，仍为正数的位置即为（未出现过）消失的数字。

func findDisappearedNumbers(nums []int) (ans []int) {
	n := len(nums)
	for _, v := range nums {
		// 当我们遍历到某个位置时，其中的数可能已经被增加过，因此需要对 n 取模来还原出它本来的值。
		v = (v - 1) % n
		//每个数都+n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return ans
}
