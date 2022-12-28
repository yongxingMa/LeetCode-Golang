package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：242
标题：有效的字母异位词
日期：2022/06/29
类型：双端队列
*/
func Test242(t *testing.T) {
	//var s = "pwwkew"
	//fmt.Println(lengthOfLongestSubstring(s))
	updateSql := fmt.Sprintf("UPDATE Alert SET isDel = 1, delComment = '%s' WHERE ID = '%d'", "已处理", 2552)
	fmt.Println(updateSql)
}

//func isAnagram(s string, t string) bool {
//	m1 := make(map[string]int)
//	m2 := make(map[string]int)
//
//	for _, v := range s {
//		m1[string(v)]++
//	}
//	for _, v := range t {
//		m2[string(v)]++
//	}
//	return reflect.DeepEqual(m1, m2)
//}

func isAnagram(s, t string) bool {
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}
