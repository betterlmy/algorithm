package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) (res [][]int) {
	// n叉树
	sort.Ints(candidates)
	// 特殊情况
	if target < candidates[0] {
		return
	}
	sum := 0
	var track []int
	backTrack(candidates, track, target, sum, &res)
	return
}

func backTrack(candidates []int, track []int, target int, sum int, res *[][]int) bool {
	if sum == target {
		// 找到目标
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return true
	}
	if sum > target {
		// 大于目标
		return true
	}
	for i := 0; i < len(candidates); i++ {
		// 选择
		track = append(track, candidates[i])
		sum += candidates[i]

		// 回溯
		ok := backTrack(candidates[i:], track, target, sum, res) // !!下一次只选择大于等于当前值的数

		// 取消选择
		track = track[:len(track)-1]
		sum -= candidates[i]
		if ok {
			// 剪枝 prune
			break
		}
	}
	return false
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
