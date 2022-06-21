package main

import (
	"fmt"
	"testing"
)

/**
序号：055
标题：跳跃游戏
日期：2022/06/20
类型：滑动窗口/中等
*/
func Test055(t *testing.T) {
	nums := []int{2, 3, 0, 0, 4}
	fmt.Println(canJump(nums))
}

/**
如果所有元素都>=1，则可直接判断为true。因为我可以一次走一步，像一只乌龟一样走到终点。
如果有元素为0，可以把0当作“坑”，为了不掉进坑里，我需要判断坑之前的位置，是否允许我一次跳多格，像一只兔子一样越过这个坑，如果可以越过这个坑，则继续往终点走，
并继续判断未来的其他坑。如果我永远都无法越过某一个坑，则返回false，我将不可能到达终点。
*/
func canJump(nums []int) bool {
	n := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= n {
			n = 1
		} else {
			n++
		}
		if i == 0 && n > 1 {
			return false
		}
	}
	return true
}
