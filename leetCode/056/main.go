package main

import (
	"fmt"
)

type MaxQueue struct {
	// 单调的队列
	queue []int
}

func makeQueue() *MaxQueue {
	return &MaxQueue{}
}
func (this *MaxQueue) push(x int) {
	lenQ := len(this.queue)
	if lenQ == 0 {
		this.queue = append(this.queue, x)
		return
	}
	if x > this.queue[lenQ-1] {
		for i := lenQ - 1; i >= 0; i-- {
			if this.queue[i] < x {
				this.queue[i] = x
			} else {
				break
			}
		}
	}
	this.queue = append(this.queue, x)
}
func (this *MaxQueue) pop() int {
	tmp := this.queue[0]
	this.queue = this.queue[1:]
	return tmp
}
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	mQueue := makeQueue()
	right := k - 1
	for i := 0; i <= right; i++ {
		mQueue.push(nums[i])
	}
	right++
	for right < len(nums) {
		res = append(res, mQueue.pop())
		mQueue.push(nums[right])
		right++
	}
	res = append(res, mQueue.pop())

	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
