package main

import "math/rand"

type RandomizedSet struct {
	nums    []int       //o(1)的随机访问 但是查找操作需要O(n), 因此不能在O(1)的时间内完成插入和删除
	indices map[int]int //O(1)的时间内完成插入和删除操作
	// 结合两个数据结构 先从maps中找到元素得到下标 根据下标从num中获取值

}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:    []int{},
		indices: map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	// 先从hashmap中查找是否存在，如果存在直接返回false，不存在则获取nums现在的长度，将新的元素插入到nums的最后，得到的len就是他的index
	if _, ok := this.indices[val]; ok {
		return false
	}
	this.indices[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	// 先从中查找是否存在，不存在则返回false，否则获取其下标，
	// 处理 slice 将nums中最后一个元素复制到这个下标中，然后删掉nums中最后一个元素
	// 处理 map 将indices中移动的元素下标值修改 并删除val这个key

	index, ok := this.indices[val] // 获取当前要删除的元素的下标
	if !ok {
		return false
	}
	lastIndex := len(this.nums) - 1         // 获取slice中最后一个下标
	this.nums[index] = this.nums[lastIndex] // 将slice中最后一个元素放入到要删除的元素的下标中
	this.nums = this.nums[:lastIndex]       // 删除slice中最后一个元素

	this.indices[this.nums[index]] = index // 修改map中刚才最后一个元素的下标
	delete(this.indices, val)              // 删除map中的val下标
	return true

}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

func main() {

}
