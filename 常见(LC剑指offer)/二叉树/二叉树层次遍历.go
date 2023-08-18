package main

// 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
func levelOrder2(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}
	needTraversal := make([]*TreeNode, 0) // 存储需要遍历的所有节点
	res := make([][]int, 0)
	needTraversal = append(needTraversal, root)
	traversal(needTraversal, &res)
	return res
}

func traversal(needTraversal []*TreeNode, res *[][]int) {
	if len(needTraversal) == 0 {
		return
	}
	// 先添加这一行的 再将下一行要遍历的放到新的数组中
	newNeedTraversal := make([]*TreeNode, 0)
	newLine := make([]int, len(needTraversal))
	for i, node := range needTraversal {
		newLine[i] = node.Val
		if node.Left != nil {
			newNeedTraversal = append(newNeedTraversal, node.Left)
		}
		if node.Right != nil {
			newNeedTraversal = append(newNeedTraversal, node.Right)
		}
	}
	*res = append(*res, newLine)
	traversal(newNeedTraversal, res)
}
