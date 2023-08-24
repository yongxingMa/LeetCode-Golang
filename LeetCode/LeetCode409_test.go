package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：409
标题：最长回文串
日期：2023/08/14
*/
func Test409(t *testing.T) {
	s := "abccccdd"
	fmt.Println(longestPalindrome1(s))
}

func longestPalindrome1(s string) int {
	mp := map[byte]int{}

	// 统计每个字符串出现的个数
	for i := 0; i < len(s); i++ {
		mp[s[i]]++
	}

	res := 0
	for _, v := range mp {
		// 如果是奇数,向下取偶数
		if v&1 == 1 {
			res += v - 1
		} else {
			res += v
		}
	}

	if res < len(s) {
		res++
	}

	return res

}
