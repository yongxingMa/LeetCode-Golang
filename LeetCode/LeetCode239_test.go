package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：239
标题：滑动窗口最大值
日期：2022/06/29
类型：双端队列
*/
func Test239(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return make([]int, 0)
	}
	//存储的 下标
	window := make([]int, 0, k)
	//存储结果
	res := make([]int, 0, len(nums)-k+1)
	for i, v := range nums {
		//i-k 表示左界
		//i>k 表示已经开始往右移动了，window[0] <= i-k 表示左边的元素可以推出了。
		if i >= k && window[0] <= i-k {
			window = window[1:len(window)]
		}
		//存储的不为空，并且新的v大于里边的值，这里是个for循环，以此比较，并弹出
		for len(window) > 0 && nums[window[len(window)-1]] < v {
			//剔除最右边的值
			window = window[0 : len(window)-1]
		}
		//存储新的下标值
		window = append(window, i)
		//拼接每一次窗口的最大值，在最左边
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}
