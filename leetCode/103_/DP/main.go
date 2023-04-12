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

type list struct {
	s []int
}

func (l *list) pop() int {
	if len(l.s) > 0 {
		tmp := l.s[0]
		l.s = l.s[1:]
		return tmp
	}
	return -1
}
func (l *list) push(v int) {
	l.s = append(l.s, v)
}

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	// dp[i]表示当amount==i最少所需的硬币数刘昂
	// base case dp[0] = 0 表示不需要
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}
	l := list{s: []int{}}
	l.push(0)
	for len(l.s) > 0 {
		i := l.pop()
		if i == -1 {
			break
		}
		// 状态转移方程
		for _, coin := range coins {
			if i+coin > amount {
				break
			}
			if dp[i+coin] != -1 {
				continue
			}
			l.push(i + coin)
			dp[i+coin] = dp[i] + 1
		}
	}

	return dp[amount]
}

func main() {
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
}
