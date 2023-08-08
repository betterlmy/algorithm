package main

import "sort"

func isStraight(nums []int) bool {
	//大小王可以是任何数字
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] < nums[i]
	})

	zeroNum := 0 // 统计王的数量
	for i := 4; i >= 0; i-- {
		if nums[i] == 0 {
			zeroNum++
		}
	}
	nums = nums[:5-zeroNum]

	for i := 0; i < len(nums)-1; i++ {
		// 判断重复
		if nums[i] == nums[i+1] {
			return false
		}
	}

	// 计算最大值和最小值的差
	diff := nums[0] - nums[len(nums)-1]
	if diff == 4 || diff == 0 {
		return true
	}
	return false

}
