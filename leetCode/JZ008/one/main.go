package main

import "fmt"

/*
给定一个含有n个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

滑动窗口
*/

func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	res := length + 1
	left, right := 0, 1

	for {
		sum := 0
		for i := left; i < right; i++ {
			// 求子串的sum
			sum += nums[i]
		}

		// 判断 如果大则尝试更新res并让左指针右移一位
		if sum >= target {
			subLength := right - left
			if subLength < res {
				res = subLength
				if res == 1 {
					return 1
				}

			}
			left++
		} else {
			right++
			if right > length {
				if res == length+1 {
					return 0
				}
				return res
			}
		}

	}
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	//nums := []int{2, 3, 1, 2, 4, 8}
	//nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(70, nums))
}
