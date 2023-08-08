package main

func pathSum(root *TreeNode, target int) [][]int {
	// 到叶子结点 所以是dfs
	res := make([][]int, 0)
	var dfs func(node *TreeNode, sum int, path []int)
	dfs = func(node *TreeNode, sum int, path []int) {
		path = append(path, node.Val)
		sum += node.Val
		if sum == target && node.Left == nil && node.Right == nil {
			res = append(res, path)
		}

		if node.Left != nil {
			dfs(node.Left, sum, path)
		}
		if node.Right != nil {
			dfs(node.Right, sum, path)
		}
	}

	dfs(root, 0, []int{})
	return res
}
