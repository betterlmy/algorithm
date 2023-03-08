package main

import "sort"

func findContentChildren(g []int, s []int) int {
	// 经典贪心算法
	// sort
	sort.Ints(g)
	sort.Ints(s)

	res := 0
	cookieCount := len(s) // 饼干的总数
	j := 0
	for _, demand := range g {
		for j < len(s) {
			sups := s[j]
			if sups >= demand {
				// 满足条件则分配一个给他
				res++         // 结果+1
				j++           // 下标+1
				cookieCount-- //剩余饼干数-1
				break         //开始判断下一个人
			}
			if j == len(s)-1 {
				// 此时这个人给到最后也不满足条件 后面的孩子要求更高 都没有份了 直接返回即可
				return res
			}
			j++
		}
	}
	return res
}
func main() {

}
