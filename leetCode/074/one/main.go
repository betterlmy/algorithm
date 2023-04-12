package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	//res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	return intervals

}
func main() {
	fmt.Println(merge([][]int{{1, 3}, {8, 10}, {2, 6}, {15, 18}}))
}
