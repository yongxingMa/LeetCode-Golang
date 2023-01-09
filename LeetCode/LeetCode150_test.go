package LeetCode

import (
	"fmt"
	"strconv"
	"testing"
)

/**
序号：150
标题：逆波兰表达式求值
日期：2023/01/09
*/
func Test150(t *testing.T) {
	var s = "the sky is blue"
	fmt.Println(reverseWord2s(s))
}

func evalRPN(tokens []string) int {
	//定义堆
	stack := []int{}
	//判断是数字，入栈，
	//如果是子符最近的两个数计算，然后入栈
	for _, token := range tokens {
		//转换成数字，如果转换失败就不是数字
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			//栈的数字减少两个
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2) // 前面的减后面的
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2) // 前面的除以后面的
			}
		}
	}
	//最后的数字就是结果
	return stack[0]
}
