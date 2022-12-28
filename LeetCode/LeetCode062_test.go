package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：62
标题：跳跃游戏
日期：2022/06/20
类型：滑动窗口/中等
*/
func Test062(t *testing.T) {

	fmt.Println(uniquePaths(51, 9))
}

/**
如果所有元素都>=1，则可直接判断为true。因为我可以一次走一步，像一只乌龟一样走到终点。
如果有元素为0，可以把0当作“坑”，为了不掉进坑里，我需要判断坑之前的位置，是否允许我一次跳多格，像一只兔子一样越过这个坑，如果可以越过这个坑，则继续往终点走，
并继续判断未来的其他坑。如果我永远都无法越过某一个坑，则返回false，我将不可能到达终点。
*/

func uniquePaths(m int, n int) int {
	//m行n列
	dp := make([][]int, m)

	//第一列的都是1
	for i := range dp {
		//n列
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	//第一行的都是1
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
