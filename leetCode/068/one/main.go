package main

import (
	"container/heap"
	"sort"
)

/*
使用堆来实现查找第k大元素
Push 和 Pop 来实现堆的插入和删除
Add 提供接口

*/

type KthLargest struct {
	sort.IntSlice // 有序数组?
	k             int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		k: k,
	}
	for _, val := range nums {
		kl.Add(val)
	}
	return kl
}

func (kl *KthLargest) Push(val interface{}) {
	kl.IntSlice = append(kl.IntSlice, val.(int))
}

func (kl *KthLargest) Pop() interface{} {
	a := kl.IntSlice
	v := a[len(a)-1]
	kl.IntSlice = a[:len(a)-1]
	return v
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}

func main() {

}
