package main

func exchange(nums []int) []int {
	// 双指针 偶数则交换 right-- 奇数不交换 left++
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]%2 == 0 {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		} else {
			left++
		}
	}
	return nums
}
