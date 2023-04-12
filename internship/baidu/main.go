package main

import (
	"fmt"
	"math"
	"sort"
)

func min(slice []int, k int) float64 {
	n := len(slice)
	// 定义dp[i][j]表示将前i个数分成j个子序列的最小平均和
	dp := make([][]float64, n+1)

	// 初始化
	for i := range dp {
		dp[i] = make([]float64, k+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0

	// 计算dp[i][j]

	/*
		包含着两种情况:
		1. 每次将最后一个数放入一段
		2. 将最后几个数放入一段
	*/
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {

			// 终止条件 j>i
			if j > i {
				break
			}

			dp[i][j] = float64(slice[i-1]) + dp[i-1][j-1] // 第一种情况

			for t := i - 1; t >= j-1; t-- {
				// 第二种情况 检查将后面的数放入到一段与第一种情况对比 将最小值放入其中

				if dp[t][j-1] > -0 {
					avg := float64(slice[i-1]) + (dp[i-1][j-1]-dp[t][j-1])/float64(i-t-1)
					if dp[i][j] < 0 || avg < dp[i][j] {
						dp[i][j] = avg
					}
				}

			}
		}
	}
	res := math.MaxFloat64
	for _, v := range dp[len(dp)-1] {
		if v == -1 {
			continue
		}
		res = min1(v, res)
	}
	return res
	// 4 3 -1 2 -1 3
}

func min1(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func main() {

	// 处理输入
	n, k := 0, 0
	fmt.Scan(&n, &k)
	oriSlice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&oriSlice[i])
	}

	// 排序
	sort.Ints(oriSlice)
	fmt.Printf("%.6f\n", min(oriSlice, k))
}
