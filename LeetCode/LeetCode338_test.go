package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：338
标题：比特位计数
日期：2022/06/19
类型：
*/
func Test338(t *testing.T) {
	var nums = []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

func countBits(n int) []int {
	res := make([]int, n+1)
	res[0] = 0
	for i := range res {
		res[i] = onesCount(i)
		//求出i的二进制，并判断有几个1
	}
	return res
}

func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}
