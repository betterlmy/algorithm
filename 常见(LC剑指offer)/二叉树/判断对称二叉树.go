package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	for {
		if root.Left == nil && root.Right == nil {
			return true
		}
		if root.Left != nil && root.Right != nil {
			if root.Left.Val != root.Right.Val {
				return false
			}
			return same2Tree(root.Left, mirror(root.Right))
		}
		return false
	}
}

func same2Tree(a *TreeNode, b *TreeNode) bool {

	if a.Val != b.Val {
		return false
	}
	// a左不存在 a左存在 a右不存在 a右存在 16种情况
	if (a.Left != nil && b.Left == nil) || (a.Right != nil && b.Right == nil) || (a.Left == nil && b.Left != nil) || (a.Right == nil && b.Right != nil) {
		// 4种分枝不对称 false
		return false
	}
	if a.Left == nil && a.Right == nil && b.Left == nil && b.Right == nil {
		// 全空 1种 true
		return true
	}
	// 分枝对称
	same1, same2 := true, true
	if a.Left != nil {
		same1 = same2Tree(a.Left, b.Left)
	}
	if a.Right != nil {
		same2 = same2Tree(a.Right, b.Right)
	}
	return same1 && same2

}

func mirror(node *TreeNode) *TreeNode {
	// 树的镜像
	if node == nil {
		return node
	} else if node.Left != nil && node.Right != nil {
		node.Left, node.Right = node.Right, node.Left
		mirror(node.Left)
		mirror(node.Right)
	} else if node.Left == nil && node.Right != nil {
		// 左空右非空
		node.Left = node.Right
		node.Right = nil
		mirror(node.Left)
	} else if node.Right == nil && node.Left != nil {
		node.Right = node.Left
		node.Left = nil
		mirror(node.Right)
	}
	return node
}
