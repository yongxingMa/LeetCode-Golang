package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：32
标题：最长的有效括号
日期：2023/08/13
类型：动态规划
*/
func Test032(t *testing.T) {
	var s = "(()())"
	fmt.Println(longestValidParentheses2(s))
}

// ()(())
/**
1 s[i]是'('，以它为结尾的子串，肯定不是有效括号子串——dp[i] = 0
2 s[i]是')'，以它为结尾的子串可能是有效子串，还需考察前一个子问题的末尾s[i-1]
	2.1 s[i-1]是'('，s[i-1]和s[i]组成一对，有效长度保底为 2，还需考察s[i-2]：
		2.1.1 s[i-2]不存在，则有效长度只有 2——dp[i] = 2
		2.1.2 s[i-2]存在，则加上以s[i-2]为结尾的最长有效长度——dp[i]=dp[i-2]+2
	2.2 s[i-1]是')'，s[i-1]和s[i]是'))'，以s[i-1]为结尾的最长有效长度为dp[i-1]，跨过这个长度（具体细节不用管，总之它最大能提供dp[i-1]长度），来看s[i-dp[i-1]-1]这个字符：
		2.2.1 s[i-dp[i-1]-1]不存在或为')'，则s[i]找不到匹配，直接gg——dp[i]=0
		2.2.2 s[i-dp[i-1]-1]是'('，与s[i]匹配，有效长度 = 2 + 跨过的dp[i-1]+ 前方的dp[i-dp[i-1]-2]。等一下，s[i-dp[i-1]-2]要存在才行！：
			s[i-dp[i-1]-2]存在，dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
			s[i-dp[i-1]-2]不存在，dp[i] = dp[i-1] + 2

*/
// 动态规划
func longestValidParentheses(s string) int {
	res := 0
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i-2 > 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1]-2 >= 0 {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		res = max(res, dp[i])
	}
	return res
}

// 栈
func longestValidParentheses2(s string) int {
	res := 0
	if len(s) == 0 {
		return 0
	}
	// 记录下标值表示有效长度
	stack := []int{}
	stack = append(stack, -1)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				//  len(stack) - 1
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}
	return res
}
