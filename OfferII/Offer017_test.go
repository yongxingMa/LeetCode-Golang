package main

import (
	"fmt"
	"testing"
)

/**
序号：017
标题：含有所有字符的最短字符串
日期：2022/12/27
*/

func Test017(t *testing.T) {
	s := "ADOBECODEBANC"
	s2 := "ABC"
	fmt.Println(minWindow(s, s2))
}

//滑动窗口
func minWindow(s string, t string) string {
	// 左指针、右指针的最小区间
	minLeft, minRight := 0, -1
	//用哈希表记录窗口t匹配数量
	ht := map[byte]int{}

	for i := 0; i < len(t); i++ {
		ht[t[i]]++
	}
	charNum := len(ht)
	for left, right := 0, 0; right < len(s); right++ { // 窗口指针
		if _, ok := ht[s[right]]; ok { // 如果是 t 中字符
			ht[s[right]]--
			if ht[s[right]] == 0 {
				charNum--
			}

			if charNum == 0 { //charNum ==0 说明已经找到字符串t
				for left < right { // 窗口左边界右移, 缩小窗口尺寸
					if v, ok := ht[s[left]]; ok { //如果找到t的字符串，判断是否小于0，如果小于0说明右边界已经有了，可以继续左移
						if v < 0 { //负数表示窗口包含多余的字符，可以继续缩小
							ht[s[left]]++
						} else { //否则，此时的窗口不能再缩小
							break
						}
					}
					left++ //如果不是t的字符串，左边界右移
				}

				// 与之前窗口比较，记录最小的窗口
				if minRight < 0 || minRight-minLeft > right-left {
					minLeft = left
					minRight = right
				}
			}
		}
	}

	return s[minLeft : minRight+1]
}
