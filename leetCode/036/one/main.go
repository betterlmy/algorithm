package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack = make([]int, 0)
	for _, v := range tokens {
		lenS := len(stack)
		switch v {
		case "+":
			stack[lenS-2] = stack[lenS-2] + stack[lenS-1]
			stack = stack[:lenS-1]
		case "-":
			stack[lenS-2] = stack[lenS-2] - stack[lenS-1]
			stack = stack[:lenS-1]
		case "*":
			stack[lenS-2] = stack[lenS-2] * stack[lenS-1]
			stack = stack[:lenS-1]
		case "/":
			stack[lenS-2] = stack[lenS-2] / stack[lenS-1]
			stack = stack[:lenS-1]
		default:
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}

	}
	return stack[0]
}

func main() {
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))

}
