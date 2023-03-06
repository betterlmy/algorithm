package sort

import "fmt"

func InsertionSort(ints []int) []int {
	length := len(ints)
	// 默认第一个元素是有序的
	for i := 1; i < length; i++ {
		// 对每个元素进行一次插入操作
		preIndex := i - 1
		current := ints[i]
		for preIndex > 0 && current < ints[preIndex] {
			// 满足条件的则进行插入 从当前元素遍历到头或者碰到比他小的元素时停止
			ints[preIndex+1] = ints[preIndex]
			preIndex--
		}
		ints[preIndex+1] = current
	}
	fmt.Println("插入排序结束")
	return ints
}
