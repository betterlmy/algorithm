package main

/*
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
*/
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	// 初始化
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	// i j 表明要访问的元素位置 k表示要匹配第几个word 返回是否能够匹配成功
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || board[i][j] != word[k] {
			// 返回False 的情况：1. 数组越界 2.不可重复访问 3.元素不匹配
			return false
		}
		if k == len(word)-1 {
			// 说明找到了元素 已经全部找到了
			return true
		}
		visited[i][j] = true // 标记为true

		// 回溯访问下一个元素 右 下 左 上
		res := dfs(i+1, j, k+1) || dfs(i, j+1, k+1) || dfs(i-1, j, k+1) || dfs(i, j-1, k+1)
		visited[i][j] = false
		return res
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				// 只要有一个true就返回true 如果没有true 则从下一个开始
				return true
			}
		}
	}
	return false
}
func main() {

}
