package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	a, b = reverseStr(a), reverseStr(b)
	if len(a) < len(b) {
		// 保证b长度小
		a, b = b, a
	}
	res := ""
	c := 0
	for i := 0; i < len(b); i++ {
		x := int(a[i] - '0')
		y := int(b[i] - '0')
		sum := x + y + c
		if sum >= 2 {
			c = 1
			sum -= 2
		} else {
			c = 0
		}
		res += strconv.Itoa(sum)
	}
	for i := len(b); i < len(a); i++ {
		x := int(a[i] - '0')
		sum := x + c
		if sum >= 2 {
			c = 1
			sum -= 2
		} else {
			c = 0
		}
		res += strconv.Itoa(sum)
	}
	if c == 1 {
		res += "1"
	}
	return reverseStr(res)

}

func reverseStr(str string) string {
	i, j := 0, len(str)-1
	strB := []byte(str)
	for i <= j {
		strB[i], strB[j] = strB[j], strB[i]
		i++
		j--
	}
	return string(strB)
}

func main() {
	fmt.Println(addBinary("1010", "1011"))

}
