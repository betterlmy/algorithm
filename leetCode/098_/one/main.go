package main

/*
	验证二叉搜索树
	https://leetcode.cn/problems/validate-binary-search-tree/
	递归法
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		// 叶子结点
		return true
	}

	if root.Left == nil {
		// 只有右子树
		if isValidBST(root.Right) {
			if root.Val < min(root.Right) {
				return true
			}
			return false
		} else {
			return false
		}

	} else if root.Right == nil {
		// 只有左子树
		if isValidBST(root.Left) {
			if root.Val > max(root.Left) {
				return true
			}
			return false
		} else {
			return false
		}

	} else {
		// 左右子树都存在
		if isValidBST(root.Right) && isValidBST(root.Left) {
			if root.Val < min(root.Right) && root.Val > max(root.Left) {
				return true
			}
			return false
		} else {
			return false
		}
	}
}

func min(root *TreeNode) int {
	if root.Left == nil {
		// 叶子结点
		return root.Val
	}
	return min(root.Left)

}
func max(root *TreeNode) int {
	if root.Right == nil {
		// 叶子结点
		return root.Val
	}
	return max(root.Right)
}
