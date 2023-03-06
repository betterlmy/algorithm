package main

//给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
// 思路: 用一个map来存储每个元素出现的次数,然后遍历map,找到出现次数为1的元素

func singleNumber(nums []int) int {
	maps := make(map[int]int, len(nums)/3+1)

	for _, v := range nums {
		maps[v]++
	}

	for i, v := range maps {
		if v == 1 {
			return i
		}
	}
	return -1

}

func main() {
	testnum1 := []int{2, 2, 3, 2}
	testnum2 := []int{0, 1, 0, 1, 0, 1, 99}
	println(singleNumber(testnum1))
	println(singleNumber(testnum2))
}
