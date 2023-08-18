package main

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	} else if root.Left == nil && root.Right != nil {
		root.Left = root.Right
		root.Right = nil
		mirrorTree(root.Left)
	} else if root.Left != nil && root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		mirrorTree(root.Right)
	} else {
		root.Left, root.Right = root.Right, root.Left
		mirrorTree(root.Left)
		mirrorTree(root.Right)
	}
	return root
}
