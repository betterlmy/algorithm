package sort

import "fmt"

func SelectionSort(ints []int) []int {
	length := len(ints)
	if length < 2 {
		return ints
	}
	for i := 0; i < length; i++ {
		// 遍历所有
		minIndex := i
		changed := false
		for j := i + 1; j < length; j++ {
			// 未排序的子集
			if ints[minIndex] > ints[j] {
				changed = true
				minIndex = j
			}
		}
		if changed {
			ints[i], ints[minIndex] = ints[minIndex], ints[i]
		}
	}
	fmt.Println("选择排序结束")
	return ints
}
