package main

import (
	"fmt"
	"testing"
)

/**
序号：003
标题：前 n 个数字二进制中 1 的个数
日期：2022/12/22
类型：简单
*/

func Test003(t *testing.T) {
	fmt.Println(countBits(2))
}

//j & (j - 1) 表示把j的二进制最右边的1变成0，循环几次说明有几个1
func countBits(n int) []int {
	res := make([]int, n+1)
	//循环便利数组中的数
	for i := 0; i <= n; i++ {
		j := i
		//每个数进行与操作，每次与操作会少1 直到变成0，同时在res的数组中记录次数
		for j != 0 {
			res[i]++
			j = j & (j - 1)
		}
	}
	return res
}
