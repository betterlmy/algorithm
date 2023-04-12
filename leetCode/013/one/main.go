package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	nums, ori []int
}

func Constructor(nums []int) Solution {
	ori := make([]int, len(nums))
	copy(ori, nums)
	return Solution{
		nums: nums,
		ori:  ori,
	}

}

func (this *Solution) Reset() []int {
	this.nums = this.ori
	return this.nums
}

func (this *Solution) Shuffle() []int {
	n := len(this.nums)
	for i := range this.nums {
		j := i + rand.Intn(n-i)
		this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
	}
	return this.nums

}

func main() {
	s := Constructor([]int{1, 5, 8})
	fmt.Println(s.Shuffle())
	fmt.Println(s.Shuffle())
	fmt.Println(s.Reset())

}
