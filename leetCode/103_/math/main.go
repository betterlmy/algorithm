package main

import (
	"fmt"
	"sort"
)

/*
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回-1。

你可以认为每种硬币的数量是无限的。

链接：https://leetcode.cn/problems/gaM7Ch
*/

func main() {
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
}

func coinChange(coins []int, amount int) int {
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})
	res := 0
	for _, coin := range coins {
		if amount < coin {
			continue
		}
		num := amount / coin
		amount -= coin * num
		res += num
	}
	if amount == 0 {
		return res
	}
	return -1

}
