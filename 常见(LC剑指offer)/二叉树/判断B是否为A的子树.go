package main

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	// 先序遍历先获得和B.Val相同的节点
	i := 0
	ANode := []*TreeNode{A}
	for i < len(ANode) {
		node := ANode[i]
		i++
		if node.Val == B.Val {
			// 得到相同的节点 开始处理
			return same(node, B)
		} else {
			// 没有遇到 则继续先序遍历
			if node.Left != nil {
				ANode = append(ANode, node.Left)
			}
			if node.Right != nil {
				ANode = append(ANode, node.Right)
			}
		}
	}
	return false
}

func same(a *TreeNode, b *TreeNode) bool {
	ANode := []*TreeNode{a}
	BNode := []*TreeNode{b}
	i := 0
	for i < len(BNode) {
		if ANode[i].Val == BNode[i].Val {
			if ANode[i].Left != nil && BNode[i].Left != nil {
				ANode = append(ANode, ANode[i].Left)
				BNode = append(BNode, BNode[i].Left)
			} else if ANode[i].Left == nil && BNode[i].Left != nil {
				// A少一个 直接false B少一个left 无所谓
				return false
			}

			if ANode[i].Right != nil && BNode[i].Right != nil {
				ANode = append(ANode, ANode[i].Right)
				BNode = append(BNode, BNode[i].Right)
			} else if ANode[i].Right == nil && BNode[i].Right != nil {
				// A少一个 直接false B少一个left 无所谓
				return false
			}
			i++
		} else {
			return false
		}
	}
	return true
}
