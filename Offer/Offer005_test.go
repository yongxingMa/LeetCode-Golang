package main

import (
	"fmt"
	"strings"
	"testing"
)

/**
序号：005
标题：替换空格
日期：2023/01/04
*/

func Test005(t *testing.T) {
	//nums := make([]int, 0)
	//words := []string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}
	words := "abcdefgh"
	fmt.Println(replaceSpace(words))
}

//解法一 双指针
func replaceSpace(s string) string {
	space := 0
	for i := range s {
		if s[i] == ' ' {
			//为 %20 预留空间
			space++
		}
	}
	var res = make([]byte, len(s)+space*2)
	var j = len(res) - 1
	//倒序遍历
	for i := len(s) - 1; i >= 0; i-- {
		//如果是空格，连续3个拼接20%
		if s[i] == ' ' {
			res[j] = '0'
			res[j-1] = '2'
			res[j-2] = '%'
			j -= 3
		} else {
			//直接赋值
			res[j] = s[i]
			j--
		}
	}
	return string(res)
}

//解法二 额外数组
func replaceSpace2(s string) string {
	//初始化数组
	var newS []byte
	// 遍历字符串
	for i := range s {
		if s[i] == ' ' {
			//用数组替换
			newS = append(newS, '%', '2', '0')
		} else {
			newS = append(newS, s[i])
		}
	}
	return string(newS)
}

//解法三 使用stringBuilder
func replaceSpace3(s string) string {
	var res strings.Builder
	for i := range s {
		if s[i] == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteByte(s[i])
		}
	}
	return res.String()
}

//解法四 偷懒法
func replaceSpace4(s string) string {
	return strings.Replace(s, " ", "20%", -1)
}
