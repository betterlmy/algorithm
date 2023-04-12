package main

/*
	验证二叉搜索树
	https://leetcode.cn/problems/validate-binary-search-tree/
	中序遍历 二叉搜索树中序遍历一定有序
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	x := make([]int, 0)
	mid(root, &x)
	last := x[0]

	for i, v := range x {
		// 遍历一遍 检查是否有序
		if i == 0 {
			continue
		}

		if last >= v {
			return false
		} else {
			last = v
		}
	}
	return true

}

func mid(root *TreeNode, x *[]int) {
	// 中序遍历
	if root.Left != nil {
		mid(root.Left, x)
	}
	*x = append(*x, root.Val)
	if root.Right != nil {
		mid(root.Right, x)
	}
}
