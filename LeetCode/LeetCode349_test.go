package LeetCode

import (
	"fmt"
	"sort"
	"testing"
)

/**
序号：349
标题：两个数组的交集
日期：2023/01/02
*/
func Test349(t *testing.T) {
	var nums1 = []int{1, 2, 2, 1}
	var nums2 = []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}

//解法一 hashmap
func intersection(nums1 []int, nums2 []int) []int {
	hashmap := make(map[int]int)
	//遍历数组1，hashmap存储
	for _, v := range nums1 {
		hashmap[v] = 1
	}
	var res []int

	// 利用count>0，实现重复值只拿一次放入返回结果中
	for _, v := range nums2 {
		//如果在数组1中存在，并且count大于0
		// count记录的是hashmap中的值 ==1 说明数组1中有，数组2还没比较，因为比较过的都会
		if count, ok := hashmap[v]; ok && count > 0 {
			//添加到结果集
			res = append(res, v)
			//拿过的值，减1
			hashmap[v]--
		}
	}
	return res
}

//解法二 排序+双指针
func intersection2(nums1 []int, nums2 []int) (res []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if res == nil || x > res[len(res)-1] {
				res = append(res, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return
}
