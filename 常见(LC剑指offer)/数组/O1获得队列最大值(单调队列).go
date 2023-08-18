package main

import "fmt"

//请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
//若队列为空，pop_front 和 max_value 需要返回 -1

type MaxQueue struct {
	queue   []int
	deQueue DeQueue
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	return this.deQueue.getMax()
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	this.deQueue.push(value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	tmp := this.queue[0]
	this.deQueue.pop(tmp)
	this.queue = this.queue[1:]
	return tmp
}

// DeQueue 单调队列 队列的队首是最大值 从左往右依次减少
type DeQueue struct {
	queue []int
}

func (d *DeQueue) push(num int) {
	// 向双向队列插入一个值 从左到右 比他大的不变 比他小的丢弃
	if len(d.queue) == 0 {
		d.queue = append(d.queue, num)
		return
	}
	i := 0
	for i < len(d.queue) {
		if d.queue[i] >= num {
			i++
		} else {
			break
		}
	}
	d.queue = d.queue[:i]
	d.queue = append(d.queue, num)
}

func (d *DeQueue) pop(num int) {
	if len(d.queue) == 0 {
		return
	}
	if num == d.queue[0] {
		d.queue = d.queue[1:]
	}
}
func (d *DeQueue) getMax() int {
	if len(d.queue) == 0 {
		return -1
	}
	return d.queue[0]
}

func main() {
	q := Constructor()
	q.Push_back(94)
	q.Push_back(16)
	q.Push_back(89)
	fmt.Println(q.Pop_front())
	q.Push_back(22)
	fmt.Println(q.Pop_front())
}
