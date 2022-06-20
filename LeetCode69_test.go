package main

import (
	"fmt"
	"testing"
)

/**
序号：69
标题：x 的平方根
日期：2022/06/02
类型：二分查找
*/
func Test69(t *testing.T) {

	fmt.Println(mySqrt(8))
}

/**

 */
func mySqrt(x int) int {
	left,right := 0,x
	ans := -1
	for left <= right {
		mid := left + (right - left)/2
		if mid * mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}