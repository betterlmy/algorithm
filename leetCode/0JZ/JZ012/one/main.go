package main

import "fmt"

func pivotIndex(nums []int) int {

	length := len(nums)
	sum := 0
	sums := make([]int, 0)

	for _, v := range nums {
		sum += v
		sums = append(sums, sum)
	}
	total := sums[length-1]

	for i, v := range nums {
		if i == 0 && v == total {
			return i
		}
		if i == length-1 && sums[i-1] == 0 {
			return i
		}
		if i != 0 && i != length-1 && 2*sums[i-1]+v == total {
			return i
		}
	}
	return -1

}
func main() {
	fmt.Println(pivotIndex([]int{-1, -1, 0, 1, 1, 0}))

}
