package main

import (
	"algorithm/sort"
	"fmt"
)

func main() {
	x := []int{5, 66, 71, 2, 11, 66}
	fmt.Println(sort.BubbleSort(x))
	fmt.Println(sort.InsertionSort(x))
	fmt.Println(sort.SelectionSort(x))
}
