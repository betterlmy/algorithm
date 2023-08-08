package main

import (
	"fmt"
	"math"
)

// BubbleSort 冒泡排序 两两交换 添加哨兵 如果没有交换则直接输出  稳定
func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		flag := true
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
				flag = false
			}
		}
		if flag {
			return nums
		}
	}
	return nums
}

// SelectSort 选择排序 每次拿出最大或者最小的与指定需要的交换位置
func SelectSort(nums []int) []int {
	minIndex := 0
	for i := 0; i < len(nums)-1; i++ {
		// 第一轮循环 遍历所有的位置 每次遍历就得到一个合适的点
		min := math.MaxInt
		flag := true
		//minIndex := i
		for j := i; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				minIndex = j
				flag = false
			}
		}
		if flag {
			break
		} else {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
	return nums
}

// InsertSort 插入排序 稳定
func InsertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		tmp := i
		for tmp > 0 && nums[tmp] < nums[tmp-1] {
			nums[tmp], nums[tmp-1] = nums[tmp-1], nums[tmp]
			tmp--
		}
	}
	return nums
}

// MergeSort 归并排序 分治 先拆分 终止条件是拆分到只有一个或者没有元素为止  然后合并 稳定
func MergeSort(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}
func mergeSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	// 拆分
	mid := start + (end-start)/2
	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)

	// 合并
	merge(nums, start, mid, end)
}

func merge(nums []int, start int, mid int, end int) {
	// 合并两个有序数组 双指针 从后往前
	left := mid
	right := end
	finishLeft := true // true代表是左边先结束
	lens := end - start + 1
	tmpNums := make([]int, lens)
	i := 0
	for left >= start && right >= mid+1 {
		i++
		if nums[left] > nums[right] {
			tmpNums[lens-i] = nums[left]
			finishLeft = true
			left--
		} else {
			tmpNums[lens-i] = nums[right]
			finishLeft = false
			right--
		}
	}
	i++
	if finishLeft {
		// 此时右边还剩下
		for right >= mid+1 {
			tmpNums[lens-i] = nums[right]
			i++
			right--
		}
	} else {
		for left >= start {
			tmpNums[lens-i] = nums[left]
			i++
			left--
		}
	}
	copy(nums[start:end+1], tmpNums)
}

// QuickSort 快速排序  分治 确定pivot位置 然后分治两个
func QuickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	pivot := nums[0]
	less, greater := make([]int, 0), make([]int, 0)
	pivorIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < pivot {
			pivorIndex++
			less = append(less, nums[i])
		} else {
			greater = append(greater, nums[i])
		}
	}
	nums[pivorIndex] = pivot
	copy(nums[:pivorIndex], less)
	less = nil
	QuickSort(nums[:pivorIndex])
	if len(greater) > 0 {
		copy(nums[pivorIndex+1:], greater)
		greater = nil
		QuickSort(nums[pivorIndex+1:])
	}
	return nums

}

// HeapSort 堆排序
func HeapSort(nums []int) []int {
	// 堆是完全二叉树! 最后一个非叶子节点的index = len/2-1
	// 大顶堆任意节点的值都大于等于子节点的值 数组模拟堆
	// 基本思想 : 1 从倒数第一个非叶子节点构建大顶堆 2. 每次拿出堆顶的元素放到最后 将堆len-1 不断调整堆

	// 1 构建大顶堆
	for i := len(nums)/2 - 1; i >= 0; i-- {
		headAdjust(nums, len(nums), i)
	}
	// 2. 拿出堆顶与最后一个交换,减少lens,同时从堆顶调整为大顶堆
	for i := 0; i < len(nums); i++ {
		nums[0], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[0]
		headAdjust(nums, len(nums)-i-1, 0)
	}
	return nums

}

func headAdjust(nums []int, length, i int) {
	// n是视野 需要处理的nums长度 i是处理哪个非叶子节点
	left, right := i*2+1, i*2+2
	largestIndex := i
	if left < length && nums[left] > nums[largestIndex] {
		largestIndex = left
	}
	if right < length && nums[right] > nums[largestIndex] {
		largestIndex = right
	}
	if largestIndex != i {
		// 说明发生变化了
		nums[i], nums[largestIndex] = nums[largestIndex], nums[i]
		headAdjust(nums, length, largestIndex) // 继续下调
	}

}

func main() {
	x := []int{1, 2, 5, 8, 2, 3, -1}
	// -1 1 2 2 3 5 8
	//fmt.Println(BubbleSort(x))
	//fmt.Println(SelectSort(x))
	//fmt.Println(InsertSort(x))
	//fmt.Println(MergeSort(x))
	//fmt.Println(QuickSort(x))
	fmt.Println(HeapSort(x))
}
