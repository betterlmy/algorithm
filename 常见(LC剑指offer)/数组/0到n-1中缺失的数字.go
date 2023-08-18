package main

import "fmt"

/*
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
输入: [0,1,2,3,4,5,6,7,9]
输出: 8
*/

func missingNumber(nums []int) int {
	// 正常情况下 On遍历得到
	// 在一般情况中 nums[i]=i 所以 需要找到第一个nums[i]!=i的情况 使用二分查找 ologn即可得到
	res := 0
	left, right := 0, len(nums)-1
	if nums[left] != left {
		return left
	}
	if nums[right] == right {
		return right + 1
	}
	find(nums, left, right, &res)
	return res
}

func find(nums []int, left int, right int, res *int) {
	if left == right {
		*res = left
		return
	}
	if left > right {
		return
	}
	mid := (right-left)/2 + left

	if nums[mid] == mid {
		// 此时说明第一个不存在的在右边
		find(nums, mid+1, right, res)
	} else {
		// 此时在左边
		if nums[mid-1] == mid-1 {
			// 终止条件
			*res = mid
			return
		}
		find(nums, left, mid-1, res)
	}
}

func main() {
	fmt.Println(missingNumber([]int{0, 1, 3}))
}
