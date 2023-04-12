package main

import (
	"fmt"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {

	sort.Slice(envelopes, func(i, j int) bool {
		// 排序 w递增 h递减
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})
	// 这样将两个条件就变成了对一个条件进行求递增子序列的问题了 lc.300

	n := len(envelopes)
	h := make([]int, n) // 高度
	for i := range h {
		h[i] = envelopes[i][1]
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1 // 初值
	}

	// 求最长的递增子序列长度
	// dp[i] 代表以i为结尾的最长递增子序列长度
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if h[j] < h[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	// 获取dp最大值
	res := 1
	for _, v := range dp {
		res = max(res, v)
	}

	return res

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 7}, {6, 4}, {2, 3}}))

}
