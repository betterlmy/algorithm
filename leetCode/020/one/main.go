package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	leftPoint, rightPoint := 0, len(s)-1
	for {
		if s[leftPoint] != s[rightPoint] {
			return false
		}
		leftPoint++
		rightPoint--
		if leftPoint >= rightPoint {
			break
		}
	}
	return true
}
func countSubstrings(s string) int {
	// 左不加 右加1
	res := 0
	for index := 0; index < len(s); index++ {
		maxOffset := min(index, len(s)-index-1)
		for offset := 0; offset <= maxOffset; offset++ {
			leftPoint := index - offset
			rightPoint := index + offset
			tmpStr1 := s[leftPoint : rightPoint+1]
			if rightPoint+2 <= len(s) {
				tmpStr2 := s[leftPoint : rightPoint+2]
				if isPalindrome(tmpStr2) {
					res++
					fmt.Println(tmpStr2)
				}
			}
			if isPalindrome(tmpStr1) {
				fmt.Println(tmpStr1)
				res++
			} else {
				break
			}

		}
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(countSubstrings("addaccad"))

}
