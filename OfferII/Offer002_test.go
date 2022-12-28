package main

import (
	"fmt"
	"strconv"
	"testing"
)

/**
序号：002
标题：二进制加法
日期：2022/12/21
类型：简单
*/
func Test002(t *testing.T) {
	fmt.Println(addBinary("11", "10"))
}

//从字符串的右端开始做加法，逢二进一
func addBinary(a string, b string) string {
	ans := ""
	var i = len(a) - 1
	var j = len(b) - 1
	//表示是否进位
	var carry = 0
	for i >= 0 || j >= 0 {
		digitA := 0
		digitB := 0
		//提示：这里的a[i] b[j] 都是ascii码，需要减去 '0' 获得数字0或1
		if i >= 0 {
			//从最右边的一位开始减 1-0=1 0-0=0
			digitA = int(a[i] - '0')
			//向左移动一位
			i--
		}
		if j >= 0 {
			digitB = int(b[j] - '0')
			//向左移动一位
			j--
		}
		// 两个数进行相加，carry表示上一轮循环的进位
		sum := digitA + digitB + carry
		// 当数字大于等于2时，产生进位，设置carry = 1 sum回归到0
		if sum >= 2 {
			carry = 1
			sum = sum - 2
		} else {
			carry = 0
		}
		//向左边进行拼接，因为我们是从右向左进行加法，所以要向左拼接
		ans = strconv.Itoa(sum) + ans
	}
	//最后的时候判断一下是不是还有进位
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}
