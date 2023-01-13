package LeetCode

/**
序号：107
标题：二叉树的层序遍历 II
日期：2023/01/12
类型：
*/

//递归 层序遍历，然后反转数组
func levelOrderBottom(root *TreeNode) [][]int {
	arr := [][]int{}

	depth := 0

	var order func(root *TreeNode, depth int)

	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		//长度等于深度，添加空数组
		if len(arr) == depth {
			arr = append(arr, []int{})
		}
		arr[depth] = append(arr[depth], root.Val)

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	order(root, depth)

	// 反转二维数组，和传统的层序遍历唯一不一样的地方，就相当于对二维数组反着输出。
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	return arr
}
