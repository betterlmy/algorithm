package main

func singleNumbers(nums []int) []int {
	// 如果两个相同的异或 则为0 两个相同的与则不变
	// 因此将所有的异或== 两个不同的异或 x  x与这两个相与得到的还是这两个值
	x := 0
	for i := range nums {
		x ^= nums[i]
	}

	res := make([]int, 0)
	for i := range nums {
		if nums[i]&x == nums[i] {
			res = append(res, nums[i])
		}
	}
	return res

}
