package LeetCode

import (
	"fmt"
	"sort"
	"testing"
)

/*
*
序号：056
标题：合并区间
日期：2023/08/12
类型：
*/
func Test056(t *testing.T) {
	nums := []int{2, 3, 0, 0, 4}
	fmt.Println(canJump(nums))
}

// 排序+双指针
func merge1(intervals [][]int) [][]int {
	// 只有一个数组，直接返回不需要合并
	if len(intervals) < 2 {
		return intervals
	}
	//每个区间的第一个数字进行比较后排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		// 当前的第一个数先比较区间里的最大的数
		if intervals[i][0] <= res[len(res)-1][1] {
			// 比较第二个数，放较大的那个
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}
