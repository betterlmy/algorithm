package main

import "fmt"

// 双指针
/*
hash表用来存储一定不包含的
*/
func checkInclusion(s1 string, s2 string) bool {
	len1 := len(s1)
	len2 := len(s2)
	if len2 < len1 {
		return false
	}

	var m1 = make(map[string]int) // 设置一个hashmap 只要两个hashmap相等即可
	for i := range s1 {
		m1[string(s1[i])]++
	}

	var m2 = make(map[string]int) // 设置一个hashmap
	leftPoint := 0
	tmpStr := s2[leftPoint : leftPoint+len1]
	for i := range tmpStr {
		m2[string(tmpStr[i])]++
	}

	for {
		equal := checkMapEqual(m1, m2)
		if equal {
			return true
		} else {
			if leftPoint+len1 > len2-1 {
				break
			}
			// 窗口滑动 每次只用修改两个值即可
			leftStr := string(s2[leftPoint])
			rightStr := string(s2[leftPoint+len1])
			m2[leftStr]--
			m2[rightStr]++
			leftPoint++
		}
	}
	return false

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
	fmt.Println(checkInclusion("xxxz", "xxxfz"))

}
