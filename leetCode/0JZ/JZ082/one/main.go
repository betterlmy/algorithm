package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) (res [][]int) {
	/*
		要求只能使用一次
	*/
	sort.Ints(candidates)
	start := 0
	sum := 0
	var track []int
	backTrack(candidates, start, sum, target, track, &res)

	return

}

func backTrack(candidates []int, start int, sum int, target int, track []int, res *[][]int) bool {
	if sum == target {
		// 满足条件
		tmp := make([]int, len(track))
		copy(tmp, track)
		if !duplicated(res, tmp) {

			*res = append(*res, tmp)
		}
		return true
	}
	if sum > target {
		// 后续都无法满足条件
		return true
	}

	for i := start; i < len(candidates); i++ {
		// 选择
		track = append(track, candidates[i])
		sum += candidates[i]
		// 回溯
		ok := backTrack(candidates, i+1, sum, target, track, res)
		// 取消选择
		track = track[:len(track)-1]
		sum -= candidates[i]
		if ok {
			// 剪枝
			break
		}
	}
	return false
}

func duplicated(res *[][]int, tmp []int) bool {
loop1:
	for _, ints := range *res {
		if len(ints) != len(tmp) {
			continue
		}
		for i := 0; i < len(tmp); i++ {
			if ints[i] != tmp[i] {
				continue loop1
			}
		}
		return true

	}
	return false

}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
