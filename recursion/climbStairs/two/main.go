package main

import "fmt"

var HashClimbedStairs = make(map[int]int)

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	if HashClimbedStairs[n] == 0 {
		HashClimbedStairs[n] = climbStairs(n-1) + climbStairs(n-2)
	}

	return HashClimbedStairs[n]

}

func main() {
	fmt.Println(climbStairs(30))
}
