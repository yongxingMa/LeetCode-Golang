package main

import (
	"fmt"
	"testing"
)

/**
序号：006
标题：排序数组中两个数字之和
日期：2022/12/22
类型：简单
*/

func Test006(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum3(nums, target))
}

//解法一 暴力法
func twoSum(numbers []int, target int) []int {
	hashmap := map[int]int{}
	for i, num := range numbers {
		ano := target - num
		//判断hashmap中是否存在，p表示hashmap中坐标，ok表示是否有值
		if p, ok := hashmap[ano]; ok {
			return []int{p, i}
		} else {
			hashmap[num] = i
		}
	}
	return nil
}

//解法二 二分查找
//首先固定第一个数，然后通过二分查找去寻找第二个数。
func twoSum2(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		//low 第一个开始
		//high 从最后一个开始
		low, high := i+1, len(numbers)-1
		for low <= high {
			//确定中间值
			mid := (high-low)/2 + low
			//如果中间值和目标值相等，直接返回第一个数和中间值的坐标
			if numbers[mid] == target-numbers[i] {
				return []int{i, mid}
				//如果中间值比目标值大，最高值变成中间值左侧的值
			} else if numbers[mid] > target-numbers[i] {
				high = mid - 1
				//如果中间值比目标值小，最小值变成中间值右侧的值
			} else {
				low = mid + 1
			}
		}
	}
	return nil
}

//解法三 双指针
//前提 递增数列适用
func twoSum3(numbers []int, target int) []int {
	//定义两个指针，分别从第一个和最后一个开始
	low, high := 0, len(numbers)-1
	//当左指针还没碰到右指针的时候
	for low < high {
		//由于是递增数组，左右指针相加，与目标值进行判断
		//等于目标值，直接返回两个指针的坐标值，大于目标值右指针左移，小于目标值左指针右移
		sum := numbers[low] + numbers[high]
		if sum == target {
			return []int{low, high}
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return nil
}
