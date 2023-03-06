package main

// 给定一个非负整数 n ，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。
/*
0 0
1 1

2 1
3 2

4 1
5 2
6 2
7 3

8 1
9 2
10
*/

import (
	"fmt"
	"math"
)

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}

	res := make([]int, 0, n+1)
	c := 1.0 // 进制
	i := 0
	for i <= n {
		if i < 2 {
			res = append(res, i)
			i++
			continue
		}
		min := int(math.Pow(2, c))   // 2
		max := int(math.Pow(2, c+1)) //4
		if i >= min && i < max {
			res = append(res, 1+res[i-min])
			i++
			if i == max {
				c++
			}
		}

	}
	return res
}

func main() {
	fmt.Println(countBits(2))
}
