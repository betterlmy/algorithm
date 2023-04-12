package main

import "fmt"

func searchInsert(nums []int, target int) int {
	maxI := len(nums) - 1
	minI := 0
	i := (maxI + minI) / 2
	if nums[0] == target {
		return 0
	}
	for maxI >= minI {
		if nums[i] == target {
			return i
		} else if nums[i] > target {
			maxI = i - 1
			i = (maxI + minI) / 2
		} else if nums[i] < target {
			minI = i + 1
			i = (maxI + minI) / 2
		}
	}
	return minI

}
func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))

}
