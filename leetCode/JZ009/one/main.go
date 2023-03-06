package main

import "fmt"

/*
给定一个正整数数组 nums和整数 k ，请找出该数组内乘积小于 k 的连续的子数组的个数。

要点:
正整数 --> 2*4>= 8 再乘任何数都是大的 先找到单个小于k的字 组成不同的子串  单独的直接拿出来 连续的再进行计算
连续子数组的个数 --> 不能用排序
小于
空串不算把

滑动窗口?
*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	res := 0

	// 将符合要求的子串添加到一个二维数组中 O(n)
	twiceDimNums := make([][]int, 0) // 用来存储需要判断的所有子串
	tmpNums := make([]int, 0)        // 单个子串的内容

	for _, num := range nums {
		// 8,5,2,4,7,9,10,5 ==> 5,2,4,7    5
		if num < k {
			tmpNums = append(tmpNums, num)
		} else {
			if len(tmpNums) > 0 {
				twiceDimNums = append(twiceDimNums, tmpNums)
			}
			tmpNums = make([]int, 0)
		}
	}
	if len(tmpNums) > 0 {
		twiceDimNums = append(twiceDimNums, tmpNums)
	}

	// 开始处理二维数组
	for _, ints := range twiceDimNums {
		if len(ints) == 1 {
			// 单个的直接加入到结果中
			res++
			continue
		}

		// 滑动窗口
		left := 0
		length := len(ints)
		for ; left < length; left++ {
			sum := 1
			right := left
			for {
				sum *= ints[right]
				if sum < k {
					// 满足条件 接着乘
					res++
					right++
					if right >= length {
						break
					}
				} else {
					break
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 0))
}
