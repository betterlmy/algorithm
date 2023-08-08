package main

import (
	"fmt"
	"math"
)

/*
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。
例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。
*/
type stack struct {
	queue []int
	len   int
}

func (s *stack) push(num int) {
	s.queue = append(s.queue, num)
	s.len++
}

func (s *stack) pop() int {
	tmp := s.queue[s.len-1]
	s.queue = s.queue[:s.len-1]
	s.len--
	return tmp
}
func (s *stack) getTop() int {
	if s.len == 0 {
		return math.MinInt
	}
	return s.queue[s.len-1]
}

func validateStackSequences(pushed []int, popped []int) bool {
	s := stack{}
	popIndex := 0
	for _, v := range pushed {
		if v != popped[popIndex] {
			// 如果入栈的和pop的不同 则分为两种情况 第一种是获取栈顶 看看和pop是不是相同 如果不相同 则入栈 popIndex不变
			// 如果相同 则pop后进行下次循环
			for popIndex < len(popped) {
				if s.getTop() != popped[popIndex] {
					s.push(v)
					break
				} else {
					s.pop()
					popIndex++
				}
			}

		} else {
			// 如果是入栈和当前pop的相同 则直接忽略
			popIndex++
		}
	}
	for popIndex < len(popped) {
		if popped[popIndex] != s.pop() {
			return false
		}
		popIndex++
	}
	return true
}

func main() {
	fmt.Println(validateStackSequences([]int{2, 1, 0}, []int{1, 2, 0}))
}
