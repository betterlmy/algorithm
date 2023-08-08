package main

type Stack []int

func (s *Stack) push(x int) {
	*s = append(*s, x)

}

func (s *Stack) pop() {
	*s = (*s)[:len(*s)-1]
}
func validateStackSequences(pushed []int, popped []int) bool {
	var s Stack
	if len(pushed) != len(popped) {
		return false
	}
	// 假设压入栈的所有数字均不相等
	pushedNum := 0 // 下一个需要压入栈的序列
	for _, v := range pushed {
		// 模拟压入栈
		pushedNum++
		s.push(v)
		if v == popped[0] {
			break
		}
	}

	for _, v := range popped {
		if len(s) == 0 {
			s.push(pushed[pushedNum])
			pushedNum++
		}
		if s[len(s)-1] == v {
			// 此时证明当前元素不需要入栈
			s.pop()
		} else {
			// 此时需要入栈 如果在剩余的push没有找到这个元素 则证明错误
			for {
				if pushedNum >= len(pushed) {
					// 此时已经没有可以放进去的了
					return false
				}
				x := pushed[pushedNum]
				s.push(x)
				pushedNum++

				if x == v {
					s.pop()
					break
				}
			}
		}
	}
	return true

}
