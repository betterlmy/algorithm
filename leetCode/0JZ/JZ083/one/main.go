package main

import (
	"fmt"
)

func permute(nums []int) (res [][]int) {
	if len(nums) <= 0 {
		return
	}

	var track []int
	visited := make([]bool, len(nums))
	backtrack(nums, track, visited, &res)
	return
}

func backtrack(nums []int, track []int, visited []bool, res *[][]int) {

	if len(track) == len(nums) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 选择
		//fmt.Println(visited)
		if visited[i] {
			continue
		}

		track = append(track, nums[i])
		visited[i] = true
		// 回溯
		backtrack(nums, track, visited, res)
		// 取消选择
		track = track[:len(track)-1]
		visited[i] = false
	}

}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
