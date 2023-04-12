package main

import "fmt"

/*
给定一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/

func subsets1(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{}) // 添加空集
	if len(nums) > 1 {
		res = append(res, nums) // 添加全集
	}

	for maxLevel := 1; maxLevel < len(nums); maxLevel++ {
		// 针对不同的长度进行处理

		tmpNums := make([]int, len(nums))
		copy(tmpNums, nums)

		visited := make([]bool, len(tmpNums)) // 控制访问
		tmpArr := make([]int, 0)
		var backtrack func(int, int)
		backtrack = func(level int, maxLevel int) {

			if level == maxLevel {
				// 终止条件
				arr := make([]int, maxLevel)
				copy(arr, tmpArr)
				res = append(res, arr)
			}

			for i := 0; i < len(tmpNums); i++ {
				// 选择
				if visited[i] {
					continue
				}

				tmpArr = append(tmpArr, tmpNums[i]) // 选择当前
				visited[i] = true
				backtrack(level+1, maxLevel) // 递归下一层

				// 取消选择
				tmpArr = tmpArr[:len(tmpArr)-1]
				visited[i] = false
			}

		}
		backtrack(0, maxLevel)
	}

	return res
}

func subsets(nums []int) [][]int {
	var res [][]int
	var backtrack func(nums []int, start int, track []int)

	backtrack = func(nums []int, start int, track []int) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp) // 路过都添加

		for i := start; i < len(nums); i++ {
			// 做选择
			track = append(track, nums[i])

			// 回溯
			backtrack(nums, i+1, track) // i+1是关键 下一次直接跳过前面所有的i

			// 取消选择
			track = track[:len(track)-1]
		}

	}
	backtrack(nums, 0, []int{})

	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
