package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	res := 1
	last := res

	for {
		res *= 2
		nowsq := res * res
		if x == nowsq {
			return res
		}
		if x > nowsq {
			last = res
		} else {
			break
		}
	}

	// 确定范围是(last,res)
	for {
		if res-last <= 1 {
			return last
		}
		now := (res + last) / 2
		nowsq := now * now
		if nowsq > x {
			res = now
		} else if nowsq == x {
			return now
		} else {
			last = now
		}
	}
	return 0
}
func main() {
	fmt.Println(mySqrt(9))

}
