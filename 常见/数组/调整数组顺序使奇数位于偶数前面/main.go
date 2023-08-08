package main

/*
输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。
如果有多对数字的和等于s，则输出任意一对即可。
*/

func twoSum(nums []int, target int) []int {
	// 左右指针 和小则左指针右移 和大泽右指针左移
	left, right := 0, len(nums)-1
	if 2*nums[left] > target || 2*nums[right] < target {
		return nil
	}
	for left < right {
		if nums[left]+nums[right] == target {
			return []int{nums[left], nums[right]}
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
