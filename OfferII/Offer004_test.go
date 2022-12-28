package main

import (
	"fmt"
	"testing"
)

/**
序号：004
标题：只出现一次的数字
日期：2022/12/22
类型：中等
*/

func Test004(t *testing.T) {
	//nums := make([]int, 0)
	nums := []int{2, 2, 3, 2}
	fmt.Println(singleNumber2(nums))
}

//解法一
//声明map，记录出现的次数
func singleNumber(nums []int) int {
	resMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		resMap[nums[i]]++
	}
	for num, occ := range resMap {
		if occ == 1 {
			return num
		}
	}
	return 0
}

//解法二
//任何三个数字的异或结果都是数字本身
//把所有数组所有数字相加，重复3次的数字，对应的二进制位是0或是3
func singleNumber2(nums []int) int {
	//一个整数是由32位的二进制组成
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			//位运算 (x >> i) & 1 得到 x 的第 i 个二进制位
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans = ans | 1<<i
		}
	}
	return int(ans)
}
