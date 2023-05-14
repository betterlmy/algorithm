package main

// 给定一个整数数组和一个整数k，找到该数组中和为k的！连续子数组！个数
// 非正整数 非有序
func subarraySum(nums []int, k int) int {
	if len(nums) == 1 {
		if nums[0] == k {
			return 1
		}
		return 0
	}
	// 滑动窗口
	res := 0
	mp := make(map[int]int)
	mp[0] = 1
	sum := 0

	for _, num := range nums {
		sum += num
		res += mp[sum-k]
		mp[sum]++
	}
	return res
}

func main() {

}
