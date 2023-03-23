package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	if len(timePoints) > 1440 {
		return 0
	}
	var res int
	timePointsInt := convertTimePoints(timePoints)
	res = 1440 - timePointsInt[len(timePointsInt)-1] + timePointsInt[0]

	for i := 0; i < len(timePointsInt)-1; i++ {
		j := i + 1
		res = min(res, timePointsInt[j]-timePointsInt[i])
	}

	return res
}
func min(a, b int) int {
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}
	if a < b {
		return a
	}
	return b
}

// convertTimePoints 将[HH:MM]转换为分钟数
func convertTimePoints(timePoints []string) []int {
	var res []int
	for _, timePoint := range timePoints {
		h := strings.Split(timePoint, ":")[0]
		m := strings.Split(timePoint, ":")[1]
		hInt, _ := strconv.Atoi(h)
		mInt, _ := strconv.Atoi(m)
		res = append(res, hInt*60+mInt)
	}

	sort.Ints(res)
	return res
}

func main() {
	fmt.Println((-1) % 5)

}
