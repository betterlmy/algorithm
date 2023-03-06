package main

import "fmt"

//给定一个已按照 升序排列的整数数组numbers ，请你从数组中找出两个数满足相加之和等于目标数target 。
//函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 0开始计数 ，所以答案数组应当满足 0<= answer[0] < answer[1] <numbers.length。
//假设数组中存在且只存在一对符合条件的数字，同时一个数字不能使用两次。

// 思路1 O(n^2) 遍历两次
// 思路2 二分查找 因为有次序 O(nlogn) 第一个for循环 遍历所有数组 第二个二分查找
// 思路3 双指针 一个Left 一个Right 相加 Left右移只会增加 Right左移只会减少 O(n)

func twoSum(numbers []int, target int) []int {

	l, r := 0, len(numbers)-1
	for {
		// 因为一定有输出 所以直接无条件循环即可
		res := numbers[l] + numbers[r] - target
		switch {
		case res == 0:
			return []int{l, r}
		case res > 0:
			r--
		case res < 0:
			l++
		}
		if l == r {
			return []int{-1, -1}
		}
	}

}

func main() {
	numbers := []int{1, 2, 4, 6, 10}
	fmt.Println(twoSum(numbers, 10))
}
