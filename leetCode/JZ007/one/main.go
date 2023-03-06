package main

import (
	"fmt"
	"sort"
)

/*
一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。

请你返回所有和为 0 且不重复的三元组。

1.三个下标都不相同(但值可能重复) 2.和为0 3.返回所有

思路: 三指针? 在上题的基础上(在一个数组中找到和为Target的两个下标) 将target进行修改O(n^2)

*/

func twoSum(numbers []int, target int) [][]int {
	res := make([][]int, 0)
	l, r := 0, len(numbers)-1

l1:
	for {
		if l == r {
			return res
		}
		// 因为一定有输出 所以直接无条件循环即可
		sum := numbers[l] + numbers[r] + target
		switch {
		case sum == 0:
			left := numbers[l]
			right := numbers[r]
			if len(res) == 0 {
				res = append(res, []int{left, right})
			} else {
				for _, knownRes := range res {
					if left == knownRes[0] || left == knownRes[1] {
						r--
						continue l1
					}
				}
				res = append(res, []int{left, right})
				r--
				break
			}

		case sum > 0:

			r--
		case sum < 0:
			l++
		}

	}
}

func threeSum(nums []int) [][]int {
	// 处理特殊情况
	var res [][]int
	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)
	//checked := make(map[int]bool) // 使用一个map来判断是否重复
	lastNum := nums[0] - 1
	for i, num := range nums {
		if i+2 >= len(nums) {
			// 长度不足 直接返回
			break
		}
		if lastNum == num {
			continue
		}
		lastNum = num

		twoRes := twoSum(nums[i+1:], num)
		if len(twoRes) != 0 {
			for _, tmpRes := range twoRes {
				res = append(res, []int{num, tmpRes[0], tmpRes[1]})
			}
		}
	}

	return res

}
func main() {
	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(threeSum(nums))

}
