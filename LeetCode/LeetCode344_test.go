package LeetCode

import (
	"testing"
)

/**
序号：344
标题：反转字符串
日期：2023/01/04
*/
func Test344(t *testing.T) {
	//var nums = []byte{"h","e","l","l","o"}
	//fmt.Println(sortedSquares(nums))
}

func reverseString(s []byte) {
	//注意是len(s)-1
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
