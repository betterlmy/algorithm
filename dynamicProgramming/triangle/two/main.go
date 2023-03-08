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

// DP算法 自上向下
func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	if length == 1 {
		// 只有一个元素时 直接输出
		return triangle[0][0]
	}

	for rowIndex := 0; rowIndex < length-1; rowIndex++ {
		// 从第一行开始 得到每个元素从顶点到这个元素的最小距离 遍历到倒数第二行之后 最后一行上的值就是从顶点到这个元素的最短距离
		row := triangle[rowIndex]
		for colIndex := range row {
			if colIndex == 0 {

				// 这一行的第一个元素 需要更新下一行的第一个(多更新的)和第二个
				triangle[rowIndex+1][0] += row[0] // 左侧元素直接加上值即可
				if rowIndex != 0 {
					triangle[rowIndex+1][colIndex+1] = min(triangle[rowIndex+1][colIndex+1]+row[colIndex], triangle[rowIndex+1][colIndex+1]+row[colIndex+1])
				} else {
					triangle[rowIndex+1][colIndex+1] = triangle[rowIndex+1][colIndex+1] + row[colIndex]
				}

			} else if colIndex == rowIndex && rowIndex != 0 {
				// 直接加上 不需要比较 因为最右侧的元素只有一条路径到达.
				triangle[rowIndex+1][colIndex+1] += row[colIndex]
			} else {
				// 除了第一个和最后一个需要特殊处理,这一行的每个元素只需要处理下一行右指针连接的那个元素 需要通过比较得到最小值,因为每个元素有两个路径到达
				//triangle[rowIndex+1][colIndex] = min(triangle[rowIndex+1][colIndex]+row[colIndex], triangle[rowIndex+1][colIndex]+row[colIndex+1])
				triangle[rowIndex+1][colIndex+1] = min(triangle[rowIndex+1][colIndex+1]+row[colIndex], triangle[rowIndex+1][colIndex+1]+row[colIndex+1])
			}
		}

	}
	lastRow := triangle[length-1]
	minLength := lastRow[0]
	for _, data := range lastRow {
		if data < minLength {
			minLength = data
		}
	}

	return minLength
}

func main() {
	triangle := [][]int{
		{-7},
		{-2, 1},
		{-5, -5, 9},
		{-4, -5, 4, 4},
		{-6, -6, 2, -1, -5},
		{3, 7, 8, -3, 7, -9},
		{-9, -1, -9, 6, 9, 0, 7},
		{-7, 0, -6, -8, 7, 1, -4, 9},
		{-3, 2, -6, -9, -7, -6, -9, 4, 0},
		{-8, -6, -3, -9, -2, -6, 7, -5, 0, 7},
		{-9, -1, -2, 4, -2, 4, 4, -1, 2, -5, 5},
		{1, 1, -6, 1, -2, -4, 4, -2, 6, -6, 0, 6},
		{-3, -3, -6, -2, -6, -2, 7, -9, -5, -7, -5, 5, 1}}
	fmt.Println(minimumTotal(triangle))
}
