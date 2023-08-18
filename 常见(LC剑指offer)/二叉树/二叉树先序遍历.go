package main

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	needTraversal := []*TreeNode{root} // 存储需要遍历的所有节点
	res := make([]int, 0)
	i := 0
	for i < len(needTraversal) {
		node := needTraversal[i]
		res = append(res, node.Val) // 添加当前节点的值
		// 添加左右子节点

		if node.Left != nil {
			needTraversal = append(needTraversal, node.Left)
		}
		if node.Right != nil {
			needTraversal = append(needTraversal, node.Right)
		}
		i++
	}
	return res
}
