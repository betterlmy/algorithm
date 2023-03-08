package backtracking

import "fmt"

var (
	cols  = make([]int, 0)
	diag1 = make([]int, 0)
	diag2 = make([]int, 0)
	/*
		对角线总共分为两种 正对角线和负对角线
		在每条正对角线上的所有元素和是相等且唯一的,即与元素(x,y)在一个正对角线上的元素包括(x+1,y-1)等等,他们的和都是x+y为固定值
		在每条正对角线上的所有元素差是相等且唯一的,即与元素(x,y)在一个负对角线上的元素包括(x+1,y+1)等等,他们的差都是x-y为固定值
		因此在对角线的slice中,每个数代表一个对角线
	*/
	n = 3
)

type board struct {
	data [][]int
}

func (b board) PrintBoard() {
	for _, rowData := range b.data {
		for _, data := range rowData {
			if data == 0 {
				fmt.Print(" *")
			} else {
				fmt.Print(" Q")
			}
		}
		fmt.Println()
	}

}

func contains(items []int, i int) bool {
	for _, item := range items {
		if item == i {
			return true
		}
	}
	return false
}
func check(row, col int, cols, diag1, diag2 []int) bool {
	if contains(cols, col) || contains(diag1, row+col) || contains(diag2, row-col) {
		// 不用判断行,因为我们是每次选择一个新行进行插入 不会有重复行出现
		return false
	}
	return true
}

func SolveNQueens(num int) (b board) {
	n = num
	b.data = make([][]int, n)

	// 初始化棋盘
	for i := range b.data {
		b.data[i] = make([]int, n)
		for j := range b.data[i] {
			b.data[i][j] = 0
		}
	}

	// 回溯求解
	backtracking(b, 0)

	return b
}

func backtracking(b board, row int) bool {
	// 回溯的进行调试
	if row == n {
		return true
	}
	rowData := b.data[row]
	for col := 0; col < len(rowData); col++ {
		if check(row, col, cols, diag1, diag2) {
			// 尝试向其中添加一个棋子  for循环只有在回溯之后才会继续进行
			b.data[row][col] = 1
			cols = append(cols, col)
			diag1 = append(diag1, row+col)
			diag2 = append(diag2, row-col)
			// 开始添加下一行 ,如果下一行返回了 直接就会return b, 如果下一行没有执行成功 则撤销刚刚的操作 继续遍历一下当前行的下一列

			if backtracking(b, row+1) {
				return true
			}
			b.data[row][col] = 0
			cols = cols[:len(cols)-1]
			diag1 = diag1[:len(diag1)-1]
			diag2 = diag2[:len(diag2)-1]

		}
	}
	return false
}
