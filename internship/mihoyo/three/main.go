package main

import (
	"fmt"
)

type slices struct {
	colors int
	length int
}

func main() {
	var lenN int
	fmt.Scan(&lenN)
	nums := make([]int, lenN)
	for j := 0; j < lenN; j++ {
		fmt.Scan(&nums[j])
	}
	s := make([]slices, 0)
	for i := 0; i < lenN; i++ {
		s = append(s, slices{
			colors: 1,
			length: 1,
		})
	}

	res := 0
	if lenN <= 2 {
		fmt.Println(lenN)
		return
	}

	for i, num := range nums {
		if i+1 >= lenN {
			break
		}
		if num == nums[i+1] {
			// 合并
			res += 1
			s[num+1].colors += s[num].colors
			s[num+1].length += s[num].length
			s[num].length = 0
			s[num].colors = 0
		}
	}

	for _, s2 := range s {
		res += s2.colors
	}

	fmt.Println(res)

}
