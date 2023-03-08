package main

import (
	"algorithm/search"
	"algorithm/sort"
	"fmt"
)

func main() {
	x := []int{5, 66, 71, 2, 11, 66}
	fmt.Println(x)
	//fmt.Println(sort.BubbleSort(x))
	//fmt.Println(sort.InsertionSort(x))
	//fmt.Println(sort.SelectionSort(x))
	// sort.MergeSort(x)
	x = sort.QuickSort(x)
	fmt.Println(x)
	fmt.Println(search.BinarySearch(x, 71))

}
