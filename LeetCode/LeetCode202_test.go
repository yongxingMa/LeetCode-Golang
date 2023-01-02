package LeetCode

import (
	"fmt"
	"testing"
)

/**
序号：202
标题：快乐数
日期：2022/01/02
*/
func Test202(t *testing.T) {
	fmt.Println(isHappy1(19))
}

//题目中说了会 无限循环，那么也就是说求和的过程中，sum会重复出现，这对解题很重要！

//解法一 哈希表判断是否循环
func isHappy1(n int) bool {
	hashmap := make(map[int]bool)
	//n不等于1 并且 哈希表中不存在，还要继续循环
	for n != 1 && !hashmap[n] { //循环简写： for ; n != 1 && !hashmap[n]; n, hashmap[n] = step(n), true { }
		n = step(n)
		hashmap[n] = true
	}
	return n == 1
}

//解法二 快慢指针求解
func isHappy2(n int) bool {
	slow, fast := n, step(n)
	//最后要不fast-1 要不 slow=fast
	for fast != 1 && slow != fast {
		//慢指针走一步
		slow = step(slow)
		//快指针走两步
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	//数位分离，先从个位数开始求平方，最后相加
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
