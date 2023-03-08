package main

import "fmt"

// 01背包问题的回溯法解
var (
	capacity = 100
	maxW     = 0
	weights  = []int{50, 10, 15, 32, 39}
	numItems = len(weights)
)

func f(i, cw int) {
	// i表示当前判断的物品  cw表示当前背包的重量
	if i == numItems {
		fmt.Println()
		// 终止条件1
		if cw > maxW {
			maxW = cw
		}
		return
	}

	f(i+1, cw) // 不加入当前物品

	newcw := cw + weights[i]
	if newcw == capacity {
		//终止条件2
		maxW = capacity
		return
	}
	if newcw < capacity {
		// 剪枝 如果加入当前物品后超出背包容量,则不进行后续任何与当前物品有关的回溯操作
		f(i+1, newcw)
	}
}

func main() {
	f(0, 0)
	fmt.Println(maxW)
}
