package main

import (
	"fmt"
	"sort"
)

// 给定一个整数数组和一个整数k，找到该数组中和为k的！连续子数组！个数
// 非正整数 非有序
// Wrong!!!错误的代码 滑动窗口无法解决非有序的案例

func subarraySum(nums []int, k int) int {
	// 滑动窗口
	res := 0
	left, right := 0, 0
	length := len(nums)

	sort.Ints(nums)

	sum := 0
	changedLeft := true // 左指针右移时重新开始求和 右指针右移累加即可
	for left <= right && right < length {
		if changedLeft {
			sum = 0
			for i := left; i < right+1; i++ {
				sum += nums[i]
			}
			changedLeft = false
		} else {
			sum += nums[right]
		}
		// 不同的值进行判断
		if sum == k {
			res += 1
			// 满足条件 right再右移则无法再相等？ 除非=0
			right++
		} else if sum < k {
			right++
			if right >= length {
				changedLeft = true
				left++
				right = left
			}
		} else {
			changedLeft = true
			left++
			right = left
		}
	}
	return res

}
func main() {
	nums := []int{1, 2, 1, 2, 1}
	k := 3
	fmt.Println(subarraySum(nums, k))

}
