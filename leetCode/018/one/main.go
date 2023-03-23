package main

import (
	"fmt"
)

// filter 过滤大小写和其他字符
func filter(s string) string {
	z := ""
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			z += string(v - ('A' - 'a'))
		} else if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			z += string(v)
		}
	}

	return z
}

func isPalindrome(s string) bool {

	s = filter(s)

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

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))

}
