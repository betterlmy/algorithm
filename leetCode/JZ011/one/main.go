package main

import "fmt"

// 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
//
// 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度
func findMaxLength(nums []int) int {

	if len(nums) == 1 {
		return 0
	}
	// //将0转为-1，求和为0的最长子串长度
	// for i,num := range nums{
	//     if num == 0{
	//         nums[i] = -1
	//     }
	// }
	maxLength := 0
	mp := make(map[int]int) // k存储值，v存储这个值第一次出现的下标
	mp[0] = -1
	counter := 0
	for i, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter--
		}
		if index, exist := mp[counter]; exist {
			maxLength = max(i-index, maxLength)
		} else {
			mp[counter] = i
		}

	}
	return maxLength

}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func main() {
	x := []int{1, 0, 1, 0, 1, 1, 1, 1, 1}
	fmt.Println(findMaxLength(x))
}
