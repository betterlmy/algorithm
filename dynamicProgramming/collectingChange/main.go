package main

import (
	"fmt"
)

/*
k种面值的硬币,分别是c1,c2,c3...ck,数量无限
给一个总金额amount 最少需要几枚硬币凑出这个金额 无法凑出则返回-1

他具有最优子结构,子问题之间相互独立
*/

func coinChange(coins []int, amount int) int {
	// 1. 确定base case, 也就是递归的终止条件
	// 2. 确定状态,也就是原问题和子问题中变化的变量,原问题是求最少需要几枚硬币凑出amount,子问题是求凑出amount-a的最少硬币数
	// 3. 确定选择,也就是导致状态发生变化的行为,选择一枚硬币
	// 4. 确定dp数组的含义,dp[i]表示凑出金额i所需的最少硬币数
	dp := make([]int, amount+1)
	dp[0] = 0
	if amount < 0 {
		return -1
	}
	for i := 1; i < amount+1; i++ {
		dp[i] = amount
	}
	for i := 1; i < amount+1; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))

}
