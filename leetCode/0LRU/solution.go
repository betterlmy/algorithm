package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	// 正常的LRU要求使用一个双向链表即可
	// 如果要求O(1)的时间找到 那就是需要借助到hash 作为缓存来存储映射关系
	cap   int        // 容量
	list  *list.List // list中存key
	cache map[int]int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		list:  list.New(),
		cache: make(map[int]int),
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.cache[key]; ok {
		// 将k移动到最后
		for e := this.list.Front(); e != nil; e = e.Next() {
			if e.Value.(int) == key {
				this.list.MoveToBack(e)
				break
			}
		}
		return val
	} else {
		return -1
	}

}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		this.cache[key] = value
		for e := this.list.Front(); e != nil; e = e.Next() {
			if e.Value.(int) == key {
				this.list.MoveToBack(e)
				break
			}
		}
	} else {
		fmt.Println("len", this.list.Len())
		if this.list.Len() == this.cap {
			delete(this.cache, this.list.Front().Value.(int))
			this.list.Remove(this.list.Front())
		}
		this.cache[key] = value
		this.list.PushBack(key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
