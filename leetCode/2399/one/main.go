package main

import "fmt"

/*
给你一个下标从 0 开始的字符串 s ，该字符串仅由小写英文字母组成，s 中的每个字母都 恰好 出现 两次 。另给你一个下标从 0 开始、长度为 26 的的整数数组 distance 。

字母表中的每个字母按从 0 到 25 依次编号（即，'a' -> 0, 'b' -> 1, 'c' -> 2, ... , 'z' -> 25）。

在一个 匀整 字符串中，第 i 个字母的两次出现之间的字母数量是 distance[i] 。如果第 i 个字母没有在 s 中出现，那么 distance[i] 可以 忽略 。

如果 s 是一个 匀整 字符串，返回 true ；否则，返回 false 。
*/
func checkDistances(s string, distance []int) bool {
	// 题目保证每个字母都出现两次

	m1 := make(map[byte]int) // m1用来存储两个字段之间的距离

	for i := 0; i < len(s); i++ {
		if v, ok := m1[s[i]]; ok {
			m1[s[i]] = i - v - 1
		} else {
			m1[s[i]] = i
		}
	}

	for i := 0; i < 26; i++ {
		if v, ok := m1[byte(i+'a')]; ok {
			if v != distance[i] {
				return false
			}
		}
	}
	return true

}

func main() {
	fmt.Println(checkDistances("abacbc", []int{1, 2, 1, 1}))

}
