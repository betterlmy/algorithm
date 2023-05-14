package main

import (
	"fmt"
)

func main() {
	/*
		灯带N个颜色随机的灯珠组成,从左往右编号为ci
		为了方便定位灯带的头和尾,头尾灯珠分别选择不同颜色的灯珠.且其他所有位置的灯珠都要与头尾灯珠的颜色不同
	*/
	// 7 1 2 3 4 3 2 5
	N := 0 // N = 灯带上灯珠的数量
	fmt.Scan(&N)
	lights := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&lights[i])
	}
	var res int64
	res = 0
	if N <= 1 {
		fmt.Println(0)
		return
	}
	if N == 2 {
		if lights[0] == lights[1] {
			fmt.Println(0)
		}
		fmt.Println(1)
		return
	}
	times := make(map[int]int)

	for _, l := range lights {
		times[l]++
	}

	left := 0
	right := len(lights) - 1
	backTrace(left, right, lights, &res, times)
	fmt.Println(res)
}

func backTrace(left int, right int, lights []int, res *int64, times map[int]int) {
	if left == right {
		return
	}

	times[lights[left]]--
	times[lights[right]]--
	if times[lights[left]] == 0 && times[lights[right]] == 0 {
		// 左右唯一
		if lights[left] != lights[right] {
			*res += 1
		}
	}
	times[lights[right]]++
	backTrace(left+1, right, lights, res, times)
	times[lights[right]]--

	times[lights[left]]++
	backTrace(left, right-1, lights, res, times)
	times[lights[left]]--
}
