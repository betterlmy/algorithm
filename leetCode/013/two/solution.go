package main

func movingCount(m int, n int, k int) int {
	// 限制只能向右或者向下走
	// 使用dfs

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var dfs func(row, col int)
	dfs = func(row, col int) {
		// 超界直接返回
		if row >= m || col >= n || visited[row][col] || getSum(row, col) > k {
			return
		}
		visited[row][col] = true
		dfs(row+1, col)
		dfs(row, col+1)
	}

	dfs(0, 0)
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] {
				res++
			}
		}
	}
	return res
}

func getSum(a, b int) int {
	sum := 0
	for a != 0 {
		x := a % 10
		sum += x
		a -= x
		a /= 10
	}
	for b != 0 {
		x := b % 10
		sum += x
		b -= x
		b /= 10
	}
	return sum
}
