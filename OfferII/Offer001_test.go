package main

import (
	"fmt"
	"math"
	"testing"
)

/**
序号：001
标题：整数除法
日期：2022/12/21
类型：/简单
*/
func Test001(t *testing.T) {
	fmt.Println(divide(15, 2))
}

func divide(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	var negative = 2
	//将所有的正数转换成负数，返回结果的时候再取反，negative判断符，=1说明是负数
	if a > 0 {
		negative--
		a = -a
	}

	if b > 0 {
		negative--
		b = -b
	}
	res := divideCore(a, b)
	if negative == 1 {
		return -res
	} else {
		return res
	}
}

//使用减法实现两个负数的除法
func divideCore(a int, b int) int {
	var res = 0
	//这里是用负数判断的，a是被除数 b是除数
	for a <= b {
		//商的值默认是1
		var quotient = 1
		value := b
		//这里的value+value是除数翻倍的意思
		//找到一个除数最大的倍数的值value，但是小于被除数的二分之一
		//找到以后跳出这个循环，被除数减去这个最大数，回到上边的那个循环
		//以15/2 为例，首选第一次循环找到value=-8 quotient=4，第二次循环找到value=-4 quotient=2，第三次循环value=-2 quotient=1
		for value >= math.MinInt32/2 && a <= value+value {
			// 如果两个除数加起来小于被除数，商+1
			quotient += quotient
			value += value
		}
		res += quotient
		a -= value
	}
	return res
}
