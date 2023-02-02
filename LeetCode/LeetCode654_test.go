package LeetCode

/**
序号：654
标题：最大二叉树
日期：2023/02/02
*/

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// 找到最大值
	index := findMax(nums)
	// 构造二叉树
	root := &TreeNode{
		Val:   nums[index],
		Left:  constructMaximumBinaryTree(nums[:index]),   //左半边
		Right: constructMaximumBinaryTree(nums[index+1:]), //右半边
	}
	return root
}
func findMax(nums []int) (index int) {
	for i, v := range nums {
		if nums[index] < v {
			index = i
		}
	}
	return
}
