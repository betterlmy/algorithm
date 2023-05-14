package main

import "fmt"

func combine(n int, k int) (res [][]int) {
	// k限制了递归树的高度
	// n是nums的长度

	// 处理特殊情况S
	if n < 0 || k < 0 || n < k {
		return
	}
	nums := make([]int, n)
	for i := 0; i < len(nums); i++ {
		nums[i] = i + 1
	}
	if n == k {
		res = append(res, nums)
		return
	}
	// 处理特殊情况E

	// 剩下的就是对决策树进行遍历了,
	backTrack(nums, k, 0, []int{}, &res)
	return
}

func backTrack(nums []int, k int, start int, trace []int, res *[][]int) {
	// 终止条件
	if len(trace) == k {
		tmp := make([]int, k)
		copy(tmp, trace)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(nums); i++ {
		// 选择
		trace = append(trace, nums[i])
		// 回溯
		backTrack(nums, k, i+1, trace, res)
		// 取消选择
		trace = trace[:len(trace)-1]
	}

}

func main() {
	fmt.Println(combine(4, 2))
}
