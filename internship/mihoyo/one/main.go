package main

import (
	"fmt"
	"sort"
)

/*
拿到一个数组,A取两个元素然后放回,B取两个然后放回,比较乘积大小 小的输掉比赛
A想知道自己有多少种选择方法是不可能输掉比赛的
数组的每个数字都是大于等于1的
8 8 8
32 23 13 12 21 31


*/

func main() {

	var lenN int
	fmt.Scan(&lenN)
	nums := make([]int, lenN)
	for j := 0; j < lenN; j++ {
		fmt.Scan(&nums[j])
	}

	sort.Slice(nums, func(i, j int) bool {
		// 排序
		return nums[i] > nums[j]
	})
	sum := 0

	last := nums[0]
	for i, num := range nums {
		if i > lenN-1 {
			break
		}
		if last == num {
			sum += 1
		} else {
			if sum >= 2 {
				break
			} else if sum < 2 {
				sum += 1
				last = num
			} else {
				break
			}
		}
	}

	// 返回A(sum,2) 以sum为底 2为指数的排列数
	if sum <= 1 {
		fmt.Println(0)
	} else {
		fmt.Println(sum * (sum - 1))
	}

}
