package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：343
标题：整数拆分
日期：2023/01/04
*/
func Test343(t *testing.T) {
	//var nums = []byte{"h","e","l","l","o"}
	fmt.Println(integerBreak(10))
}

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max(curMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}
	return dp[n]
}
