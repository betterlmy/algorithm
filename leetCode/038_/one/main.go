package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	//
	lens := len(temperatures)
	stack := make([]int, 0)
	res := make([]int, lens)
	for i, v := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, i) // 栈中存储着下标,表明下标这一天还没有出现比他温度更高的一天
			continue
		}
		for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
			// 当前元素的值 大于栈顶的元素所对应的下标的元素值
			// 说明现在的问题比存储时的温度要高 更新一下时间差值
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return res
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
