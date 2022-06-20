package main

import (
	"fmt"
	"testing"
)

/**
序号：367
标题：有效的完全平方数
日期：2022/06/02
类型：二分查找
*/
func Test367(t *testing.T) {

	fmt.Println(isPerfectSquare(7))
}

/**

 */
func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		//mid中间值坐标
		mid := (right-left)/2 + left
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid -1
		} else {
			left = mid + 1
		}
	}
	return false
}
