package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 3
	}
	if n == 4 {
		return 5
	}
	if n == 5 {
		return 8
	}
	if n == 6 {
		return 13
	}
	if n == 7 {
		return 21
	}
	if n == 8 {
		return 34
	}

	if n > 8 {
		return climbStairs(n-1) + climbStairs(n-2)
	}
	return 0

}

func main() {
	fmt.Println(climbStairs(30))
}
