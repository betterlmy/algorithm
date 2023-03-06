package sort

import "fmt"

func BubbleSort(ints []int) []int {
	length := len(ints)
	if length <= 1 {
		return ints
	}
	for i := 0; i < length; i++ {
		// 外层循环控制轮数
		for j := 0; j < length-i-1; j++ {
			// 内层循环控制每轮比较次数
			if ints[j] > ints[j+1] {
				ints[j], ints[j+1] = ints[j+1], ints[j]
			}
		}
	}
	fmt.Println("冒泡排序结束")
	return ints
}
