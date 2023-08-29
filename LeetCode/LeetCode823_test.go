package LeetCode

import (
	"fmt"
	"sort"
	"testing"
)

/*
*
序号：823
标题：带因子的二叉树
日期：2023/08/29
*/
func Test823(t *testing.T) {
	fmt.Println(findAnagrams("aa", "bb"))
}

// 考虑以数x为根进行二叉树构造的方案，然后就得到了子问题，数x的左子树的构造方案以及其右子树的构造方案，那么提前哈希表记录然后记搜即可！
func numFactoredBinaryTrees(arr []int) int {
	// 将 arr 从小到大进行排序
	// 对于i的跟节点，子节点的区间 [0,i-1)
	sort.Ints(arr)
	// dp数组保存根为i的子结点点树
	dp := make([]int64, len(arr))
	// 返回 对 1e9 + 7 取余 的结果。
	res, mod := int64(0), int64(1e9+7)
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		for left, right := 0, i-1; left <= right; left++ {
			for left <= right && int64(arr[left])*int64(arr[right]) > int64(arr[i]) {
				right--
			}
			// 满足每个非叶结点的值应等于它的两个子结点的值的乘积。
			if left <= right && int64(arr[left])*int64(arr[right]) == int64(arr[i]) {
				// left != right 时候，两个子节点可以交换所以是两种X2
				if left == right {
					dp[i] = (dp[i] + dp[left]*dp[right]) % mod
				} else {
					dp[i] = (dp[i] + dp[left]*dp[right]*2) % mod
				}
			}
		}
		res = (res + dp[i]) % mod
	}
	return int(res)
}
