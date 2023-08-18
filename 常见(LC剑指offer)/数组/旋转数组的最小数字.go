package main

import "fmt"

/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
给你一个可能存在**重复**元素值的数组 numbers ，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。
请返回旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，该数组的最小值为 1。
输入：numbers = [3,4,5,3,3]
输出：3
*/
func minArray(numbers []int) int {
	// 二分找到一个比首个元素要小的值 然后向左遍历
	first := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] != first {
			// 删除后续连续与first相同的值
			tmp := []int{first}
			numbers = append(tmp, numbers[i:]...)
			break
		}
	}

	// 向右二分查找 找到一个小于等于first的值
	left, right := 1, len(numbers)-1
	res := first
	bSearch(numbers, left, right, first, &res)
	return res
}

func bSearch(numbers []int, left, right, first int, res *int) {
	if left > right {
		return
	}
	mid := (right-left)/2 + left
	if numbers[mid] > first {
		// res还在右边
		bSearch(numbers, mid+1, right, first, res)
	} else {
		*res = numbers[mid]
		bSearch(numbers, left, mid-1, first, res)
	}
}

func main() {
	fmt.Println(minArray([]int{1, 3, 5}))
}
