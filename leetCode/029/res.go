package main

func spiralOrder(matrix [][]int) []int {
	// res := make([]int,0)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	return Traversal(matrix)
}

func Traversal(matrix [][]int) []int {
	var res []int
	rowNum := len(matrix)
	colNum := len(matrix[0])
	if rowNum == 1 {
		res = append(res, matrix[0]...)
		return res
	}
	if colNum == 1 {
		for i := 0; i < rowNum-1; i++ {
			res = append(res, matrix[i][0])
		}
		return res
	}

	res = append(res, matrix[0]...) // 添加第一行
	for i := 1; i < rowNum; i++ {
		// 添加最右侧
		res = append(res, matrix[i][colNum-1])
	}
	for i := colNum - 2; i >= 0; i-- {
		// 添加最下面
		res = append(res, matrix[rowNum-1][i])
	}
	for i := rowNum - 2; i >= 1; i-- {
		// 添加最左侧
		res = append(res, matrix[i][0])
	}
	if rowNum == 2 || colNum == 2 {
		return res
	}
	matrix = matrix[1:]
	matrix = matrix[:len(matrix)-1]
	for i := range matrix {
		matrix[i] = matrix[i][1 : len(matrix[i])-1]
	}
	res = append(res, Traversal(matrix)...)
	return res
}
