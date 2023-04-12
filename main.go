package main

import (
	"strings"
)

func maskPII(s string) string {
	res := ""
	at := strings.Index(s, "@")
	if at > 0 {
		s = strings.ToLower(s)
		res = string(s[0]) + "*****" + string(s[at-1]) + s[at:]

	} else {
		numStr := ""
		for _, v := range s {
			if v >= '0' && v <= '9' {
				numStr += string(v)
			}
		}
		s = numStr
		lens := len(s)

		pres := []string{"", "+*-", "+**-", "+***-"}

		res = pres[lens-10] + "***-***-" + s[lens-4:]
	}

	return res

}
func main() {
	//x := []int{5, 66, 71, 2, 11, 66}
	//fmt.Println(x)
	//fmt.Println(sort.BubbleSort(x))
	//fmt.Println(sort.InsertionSort(x))
	//fmt.Println(sort.SelectionSort(x))
	// sort.MergeSort(x)
	//x = sort.QuickSort(x)
	//fmt.Println(x)
	//fmt.Println(search.BinarySearch(x, 71))
	//board := backtracking.SolveNQueens(8)
	//board.PrintBoard()
	maskPII("86-(10)12345678")

}
