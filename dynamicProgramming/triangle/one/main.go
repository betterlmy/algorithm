package main

import "fmt"

/*
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
*/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// DP算法 自底向上来?
func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	if length == 1 {
		return triangle[0][0]
	}

	for i := length - 2; i >= 0; i-- {
		row := triangle[i]
		nextRow := triangle[i+1]
		for j, data := range row {
			row[j] = min(data+nextRow[j], data+nextRow[j+1])
		}
	}
	return triangle[0][0]

}

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}
