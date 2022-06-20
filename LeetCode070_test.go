package main

import (
	"fmt"
	"testing"
)

/**
序号：70
标题：爬楼梯
日期：2022/06/11
类型：
*/
func Test070(t *testing.T) {

	fmt.Println(climbStairs(8))
}

/**
动态规划可以用递归或迭代的方法去实现
*/
var store = make(map[int]int)

//递归的解法，自顶向下的
func climbStairs(n int) int {
	if store[n] != 0 {
		return store[n]
	}
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		res := climbStairs(n-1) + climbStairs(n-2)
		store[n] = res
		return res
	}
}

//动态规划，自底向上
func climbStairs2(n int) int {
	p, q, sum := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = sum
		sum = p + q
	}
	return sum
}
