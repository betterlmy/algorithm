package main

import "fmt"

type NumMatrix struct {
	Sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {

	sums := make([][]int, len(matrix))

	for i, rows := range matrix {
		tmpSums := make([]int, 0)
		sum := 0
		for _, item := range rows {
			sum += item
			tmpSums = append(tmpSums, sum)
		}
		sums[i] = tmpSums
	}
	return NumMatrix{sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i < row2+1; i++ {
		sum += this.Sums[i][col2]
		if col1 >= 1 {
			sum -= this.Sums[i][col1-1]
		}

	}
	return sum
}

func main() {
	matrix := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
	//matrix := Constructor([][]int{{-1}})

	fmt.Println(matrix.SumRegion(2, 1, 4, 3))
}
