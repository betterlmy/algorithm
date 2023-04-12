package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 2, -1, 3}
	k := 3
	ans := findMinAverage(nums, k)
	fmt.Println(ans)
}

func findMinAverage(nums []int, k int) float64 {
	n := len(nums)
	sort.Ints(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	target := float64(sum) / float64(k)
	ans := 0.0
	for i := 0; i < k; i++ {
		curSum := 0
		cnt := 1
		for j := n - 1; j >= 0; j-- {
			if float64(curSum+nums[j]) > target {
				break
			}
			curSum += nums[j]
			cnt++
		}
		ans += float64(curSum) / float64(cnt)
		n -= cnt
	}
	return ans
}
