package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：509
标题：斐波那契数
日期：2022/06/30
类型：
*/

func Test509(t *testing.T) {
	//s := "pwwkew"
	fmt.Println(fib2(5))
}

//动态规划
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

//dp数组

func fib2(n int) int {
	if n < 2 {
		return n
	}
	p, q, sum := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = sum
		sum = p + q
	}
	return sum
}
