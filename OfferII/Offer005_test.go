package main

import (
	"fmt"
	"testing"
)

/**
序号：005
标题：单词长度的最大乘积
日期：2022/12/22
类型：中等
*/

func Test005(t *testing.T) {
	//nums := make([]int, 0)
	//words := []string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}
	words := []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}
	fmt.Println(maxProduct(words))
}

//解法一
func maxProduct(words []string) int {
	//golang 二维数组初始化
	flags := make([][]int, len(words))
	for i := range flags {
		flags[i] = make([]int, 26)
	}

	//遍历数组中的每个单词
	for i := 0; i < len(words); i++ {
		//遍历每个单词的每个字母
		for j := 0; j < len(words[i]); j++ {
			//对出现的字母记录为1
			flags[i][words[i][j]-'a'] = 1
		}
	}
	var res = 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			var k = 0
			for k < 26 {
				//判断26个字母，在两个单词中是否同时出现，如果同时出现返回0
				if flags[i][k] == 1 && flags[j][k] == 1 {
					break
				}
				k++
			}
			//每个单词都遍历完成后，记录最大值
			if k == 26 {
				prod := len(words[i]) * len(words[j])
				res = max(prod, res)
			}
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//利用位掩码判断两个字符串是否有相同字符
func maxProduct2(words []string) (ans int) {
	masks := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			//这个什么意思？
			masks[i] |= 1 << (ch - 'a')
		}
	}

	for i, x := range masks {
		for j, y := range masks[:i] {
			if x&y == 0 && len(words[i])*len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return
}
