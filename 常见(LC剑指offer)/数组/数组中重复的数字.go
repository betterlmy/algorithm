package main

/*
找出数组中重复的数字。
在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3
*/
func findRepeatNumber(nums []int) int {
	// 1.排序 O(nlogn)
	// 2.hash
	exist := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := exist[v]; ok {
			return v
		}
		exist[v] = struct{}{}
	}
	return 0
}

func main() {

}
