package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	lens := len(s)
	if lens == 0 || lens == 1 {
		return lens
	}
	res := 1

	leftPoint := 0
	rightPoint := 1

	for {

		tmpStr := s[leftPoint:rightPoint]
		index := firstRepeatIndex(tmpStr)
		if index != -1 {
			res = max(res, len(tmpStr)-1)
			leftPoint += index + 1 // 得到这个序列中第一次出现重复的位置
			rightPoint = leftPoint + res + 1
			if rightPoint > len(s) {
				// 此时已经无法超过最大值了
				break
			}
		} else {
			rightPoint++
			if rightPoint > len(s) {
				res = max(res, len(tmpStr))
				break
			}
		}

	}
	return res

}
func firstRepeatIndex(str string) int {
	m1 := make(map[int32]int) //map需要用make方法来初始化 否则为nil无法写入还有slice channel
	for i, v := range str {
		if m1[v] == 0 {
			//第一次出现
			m1[v] = i + 1
		} else {
			//第二次出现
			return m1[v] - 1
		}
	}
	return -1

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcdef"))
}
