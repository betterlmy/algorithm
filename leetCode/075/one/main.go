package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {

	mapSlice := make(map[int]int, 0)
	/*
		mapSlice 存着arr2中每个变量的下标顺序
		[2,1,4,3,9,6]
		[0,1,2,3,4,5]
	*/
	for i, v := range arr2 {
		mapSlice[v] = i
	}

	map2 := make(map[int][]int, 0)
	/*
		map2 根据arr2的顺序 存储相应的值 最终的结果就是将map2按顺序合并
	*/
	arr3 := make([]int, 0) // arr3放着不在arr2中的元素
	for _, v := range arr1 {
		if _, ok := mapSlice[v]; ok {
			map2[v] = append(map2[v], v)
		} else {
			arr3 = append(arr3, v)
		}
	}

	res := make([]int, 0)
	for _, v := range arr2 {
		res = append(res, map2[v]...)
	}
	//sort.Ints(arr3)
	sort.Slice(arr3, func(i, j int) bool {
		return arr3[i] <= arr3[j]
	})
	res = append(res, arr3...)
	return res
}
func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
}
