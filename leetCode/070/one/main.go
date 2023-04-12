package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	// 正常情况下连续的偶数和奇数是一样的,找到一个x,x+1不相同的元素，返回x的下标即可
	if len(nums) == 1 {
		return 0
	}
	l, r := 0, len(nums)-1
	now := 0
	for l < r {
		now = (l + r) / 2
		if now == l {
			return r
		}
		if now%2 == 1 {
			// 奇数
			if nums[now] != nums[now-1] {
				r = now - 1
			} else {
				l = now
			}
		} else {
			if nums[now] != nums[now+1] {
				r = now
			} else {
				l = now + 1
			}
		}

	}
	return now

}
func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2}))

}
