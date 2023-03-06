package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	res := 0
	prepre := 1
	pre := 2
	for i := 3; i <= n; i++ {
		res = pre + prepre
		prepre = pre
		pre = res
	}

	return res

}

func main() {
	fmt.Println(climbStairs(30))
}
