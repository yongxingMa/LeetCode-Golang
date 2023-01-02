package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：242
标题：有效的字母异位词
日期：2022/06/29
*/
func Test242(t *testing.T) {
	//var s = "pwwkew"
	//fmt.Println(lengthOfLongestSubstring(s))
	updateSql := fmt.Sprintf("UPDATE Alert SET isDel = 1, delComment = '%s' WHERE ID = '%d'", "已处理", 2552)
	fmt.Println(updateSql)
}

//使用hashmap记录字母出现的次数
func isAnagram(s, t string) bool {
	record := [26]int{}
	for _, ch := range s {
		record[ch-'a']++
	}
	for _, ch := range t {
		record[ch-'a']--
	}
	return record == [26]int{}
}
