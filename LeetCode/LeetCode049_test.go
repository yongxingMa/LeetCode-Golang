package LeetCode

import (
	"fmt"
	"sort"
	"testing"
)

/*
*
序号：49
标题：字母异位词分组
日期：2023/08/14
*/
func Test049(t *testing.T) {
	var height = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(height, 0))
}

// 排序后比较
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
