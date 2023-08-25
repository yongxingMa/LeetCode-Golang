package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：438
标题：找到字符串中所有字母异位词
日期：2023/08/25
*/
func Test438(t *testing.T) {
	var height = []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(height))
}

//双指针构建滑动窗口原理：
//1.右指针right先向右滑动并在window中存储对应字符出现次数
//2.当左右指针间的字符数量（包括左右指针位置的字符）与p长度相同时开始比较
//3.比较完成后，左右指针均向右滑动一位，再次比较
//4.以后一直重复2、3，直到end指针走到字符串s的末尾

func findAnagrams(s string, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}
	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}
	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}
