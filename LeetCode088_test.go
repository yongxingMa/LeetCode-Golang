package main

import (
	"fmt"
	"testing"
)

/**
序号：088
标题：合并两个有序数组
日期：2022/06/11
类型：/简单
*/
func Test088(t *testing.T) {
	var nums = []int{3, 3}
	fmt.Println(twoSum001(nums, 6))
}

//双指针解法
//func merge(nums1 []int, m int, nums2 []int, n int) {
//	//临时数组
//	sorted := make([]int, 0, m+n)
//	//双指针
//	p1, p2 := 0, 0
//	for {
//		//p1已经到了最后一个，把nums2的全加进来
//		if p1 == m {
//			sorted = append(sorted, nums2[p2:]...)
//			break
//		}
//		//p2已经到了最后一个，把nums1的全加进来
//		if p2 == n {
//			sorted = append(sorted, nums1[p1:]...)
//			break
//		}
//		if nums1[p1] < nums2[p2] {
//			sorted = append(sorted, nums1[p1])
//			p1++
//		}
//		if nums1[p1] >= nums2[p2] {
//			sorted = append(sorted, nums2[p2])
//			p2++
//		}
//	}
//	copy(nums1, sorted)
//}

//逆向双指针
//空间复杂度是1
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, tail := m-1, n-1, m+n-1
	for ; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}
