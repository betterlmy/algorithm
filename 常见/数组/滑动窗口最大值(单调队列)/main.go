package main

import "fmt"

// 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

// deQueue 单调队列 队列的队首是最大值 从左往右依次减少
type deQueue struct {
	queue []int
}

func (d *deQueue) push(num int) {
	// 向双向队列插入一个值 从左到右 比他大的不变 比他小的丢弃
	if len(d.queue) == 0 {
		d.queue = append(d.queue, num)
		return
	}
	i := 0
	for i < len(d.queue) {
		if d.queue[i] >= num {
			i++
		} else {
			break
		}
	}
	d.queue = d.queue[:i]
	d.queue = append(d.queue, num)
}

func (d *deQueue) pop() {
	d.queue = d.queue[1:]
}
func (d *deQueue) getMax() int {
	return d.queue[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	// 维护一个单调队列
	if len(nums) < k {
		return nil
	}
	res := make([]int, len(nums)-k+1)
	q := deQueue{queue: make([]int, 0)}
	for i := 0; i < k; i++ {
		q.push(nums[i])
	}
	for i := 0; i < len(nums)-k; i++ {
		res[i] = q.getMax()
		if res[i] == nums[i] {
			q.pop()
		}
		q.push(nums[i+k])
	}
	return res
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
