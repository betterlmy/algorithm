package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) (res [][]int) {
	track := make([]int, 0)
	visited := make([]bool, len(nums))
	sort.Ints(nums)
	backtrack(nums, track, &visited, &res)
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return
}

func backtrack(nums []int, track []int, visited *[]bool, res *[][]int) {
	if len(track) == len(nums) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 选择
		if (*visited)[i] || i > 0 && !(*visited)[i-1] && nums[i-1] == nums[i] {
			// !!!多个1 当第一个1已经回溯完毕之后再开始回溯第二个1 此时直接剪枝即可
			continue
		}
		track = append(track, nums[i])
		(*visited)[i] = true

		backtrack(nums, track, visited, res)

		track = track[:len(track)-1]
		(*visited)[i] = false
	}
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 3}))
}
