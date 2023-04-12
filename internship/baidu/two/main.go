package main

import "fmt"

func huiwen(x int) string {
	// 构造一个仅由"r","e","d"构成的字符串 ,其中回文子串的数量恰好等于x
	if x == 3 {
		return "red"
	}
	if x == 2 {
		return "re"
	}
	if x == 1 {
		return "r"
	}
	if x <= 0 {
		return ""
	}

	// 递归
	return "red" + huiwen(x-3)

}
func main() {

	// 处理输入
	x := 0
	fmt.Scan(&x)

	fmt.Println(huiwen(x))

}
