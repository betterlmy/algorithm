package main

import "fmt"

func findAnagrams(s string, p string) []int {
	var res []int
	len1 := len(s)
	len2 := len(p)
	if len1 < len2 {
		return nil
	}

	var m1 = make(map[string]int) // 设置一个hashmap 只要两个hashmap相等即可
	for i := range p {
		m1[string(p[i])]++
	}

	var m2 = make(map[string]int) // 设置一个hashmap
	leftPoint := 0
	tmpStr := s[leftPoint : leftPoint+len2]
	for i := range tmpStr {
		m2[string(tmpStr[i])]++
	}

	// 第一步找到第一个相同的
	for {
		equal := checkMapEqual(m1, m2)
		if equal {
			res = append(res, leftPoint)
		}
		if leftPoint+len2 > len1-1 {
			return res
		}
		// 窗口滑动 每次只用修改两个值即可
		leftStr := string(s[leftPoint])
		rightStr := string(s[leftPoint+len2])
		m2[leftStr]--
		m2[rightStr]++
		leftPoint++

	}

}

// checkMapEqual 判断两个map的内容是否相同
func checkMapEqual(m1 map[string]int, m2 map[string]int) bool {
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))

}
