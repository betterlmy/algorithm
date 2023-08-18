package main

/*
在一个 n * m 的二维数组中，每一行都按照从左到右非递减的顺序排序，每一列都按照从上到下非递减的顺序排序。
请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
[
	[1,   4,  7, 11, 15],
	[2,   5,  8, 12, 19],
	[3,   6,  9, 16, 22],
	[10, 13, 14, 17, 24],
	[18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。
给定 target = 20，返回 false。
*/

func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 从右下角开始  如果tatget比他小 一定在第二象限里面 进行剪枝
	if len(matrix) == 0 {
		return false
	}
	rowIndex, colIndex := len(matrix)-1, len(matrix[0])-1
	for rowIndex >= 0 && colIndex >= 0 {
		if matrix[rowIndex][colIndex] < target {
			// 对角线都小 说明不存在
			return false
		}
		if matrix[rowIndex][colIndex] == target {
			// 说明是对角线的元素
			return true
		}
		if matrix[rowIndex][colIndex] > target {
			// 说明在这一行或者这一列或者第二象限
			for i := 0; i < rowIndex; i++ {
				// 检查这一列是否存在
				if matrix[i][colIndex] == target {
					return true
				}
			}
			for i := 0; i < colIndex; i++ {
				// 检查这一行是否存在
				if matrix[rowIndex][i] == target {
					return true
				}
			}
			colIndex--
			rowIndex--
		}
	}
	return false
}

func main() {

}
