package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：215
标题：数组中第k个最大元素
日期：2023/08/12
*/
func Test215(t *testing.T) {
	var height = []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen1(7, height))
}

//解法一 基于快排

//解法二 基于堆排
