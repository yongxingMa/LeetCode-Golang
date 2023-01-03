package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：383
标题：赎金信
日期：2023/01/03
*/
func Test383(t *testing.T) {
	ransomNote := "aa"
	magazine := "ab"
	fmt.Println(canConstruct(ransomNote, magazine))
}

// 判断 ransomNote 能不能由 magazine 里面的字符构成。
// magazine 中的每个字符只能在 ransomNote 中使用一次。
func canConstruct(ransomNote string, magazine string) bool {

	if len(ransomNote) > len(magazine) {
		return false
	}
	hashmap := map[string]int{}
	for i := range magazine {
		hashmap[string(magazine[i]-'a')]++
	}
	for i := range ransomNote {
		//判断是否有当前的字母
		if value, ok := hashmap[string(ransomNote[i]-'a')]; ok && value > 0 {
			hashmap[string(ransomNote[i]-'a')]--
		} else {
			return false
		}
	}
	return true
}
