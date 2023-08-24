package LeetCode

import (
	"fmt"
	"testing"
)

/*
*
序号：167
标题：多数元素
日期：2022/06/03
类型：
*/
func Test169(t *testing.T) {
	var nums = []int{2, 2, 1, 1, 3, 3, 3, 1, 1, 1, 1}
	fmt.Println(majorityElement(nums))
}

// 方法一 hashmap
// 方法二 数组排序
// 方法三 摩尔投票法
// 如果我们把众数记为 +1，把其他数记为 −1，将它们全部加起来，显然和大于 0，从结果本身我们可以看出众数比其他数多。
func majorityElement(nums []int) int {
	count := 0
	candidate := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count += 1
		} else {
			count -= 1
		}

	}
	return candidate
}
