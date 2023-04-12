package main

import "fmt"

func fb(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	// 自底向上 实现备忘录DP

	dp := make([]int, n+1) // dp table
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		// 状态转移方程 i时刻的状态是由i-1和i-2转移而来的
		// 状态转移方程也就是暴力算法
	}
	return dp[n]

}

func main() {
	fmt.Println(fb(5))

}
