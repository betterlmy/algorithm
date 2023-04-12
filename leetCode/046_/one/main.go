package main

import "fmt"

func permute(nums []int) [][]int {
	res := make([][]int, 0) //
	track := make([]int, 0) // 一条路径
	visited := make([]bool, len(nums))

	var backtrack func(int)
	backtrack = func(level int) {
		if level == len(nums) {
			// 证明此时已经遍历完所有的边了
			fmt.Println("遍历完成一次")
			tmp := make([]int, len(nums))
			copy(tmp, track)
			res = append(res, tmp)
		}

		for i := 0; i < len(nums); i++ {
			// 做选择,顺序判断是否已经不在路径中了
			if visited[i] {
				continue
			}
			fmt.Println("选择了:", nums[i])
			track = append(track, nums[i]) // 选择当前这个点
			visited[i] = true

			backtrack(level + 1) // 下一层的决策

			// 取消选择
			track = track[:len(track)-1] // 取消选择
			fmt.Println("取消选择了:", nums[i])
			visited[i] = false
		}
	}
	backtrack(0)
	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
