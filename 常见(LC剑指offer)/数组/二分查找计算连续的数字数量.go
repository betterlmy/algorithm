package main

import "fmt"

// 统计一个数字在排序数组中出现的次数。
// 输入: nums = [5,7,7,8,8,10], target = 8
// 输出: 2
func search(nums []int, target int) int {
	// 二分查找
	left, right := 0, len(nums)-1
	res := 0
	binarySearch(nums, target, left, right, &res)
	return res
}

func binarySearch(nums []int, target int, left int, right int, res *int) {

	if left < 0 || right > len(nums) || left > right || nums[left] > target || nums[right] < target {
		// 特殊情况
		return
	}
	// 先找到相同的第一个
	mid := (right-left)/2 + left
	if nums[mid] == target {
		*res++
		// 向左查找连续
		leftSearch := mid - 1
		for leftSearch >= 0 {
			if nums[leftSearch] == target {
				*res++
				leftSearch--
			} else {
				break
			}
		}
		rightSearch := mid + 1
		for rightSearch < len(nums) {
			if nums[rightSearch] == target {
				*res++
				rightSearch++
			} else {
				break
			}
		}
	} else if nums[mid] < target {
		// 剪枝
		binarySearch(nums, target, mid+1, right, res)
	} else {
		// 剪枝
		binarySearch(nums, target, left, mid-1, res)
	}
}

func main() {
	fmt.Println(search([]int{1, 1, 2}, 1))
}
