package main

import (
	"fmt"
)

var deleteOne bool

// validPalindrome 判断删除一个是否仍然回文
func validPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	leftPoint, rightPoint := 0, len(s)-1
	for {
		if s[leftPoint] != s[rightPoint] {
			if deleteOne {
				return false
			} else {
				deleteOne = true
				s1 := s[leftPoint+1 : rightPoint+1] //删除左边
				s2 := s[leftPoint:rightPoint]       // 删除右边
				if validPalindrome(s1) {
					return true
				} else {
					return validPalindrome(s2)
				}
			}
		}

		leftPoint++
		rightPoint--
		if leftPoint >= rightPoint {
			break
		}
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("deeee"))

}
