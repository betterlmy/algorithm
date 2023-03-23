package main

import "time"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	sentinelA := headA
	sentinelB := headB
	changedA, changedB := false, false
	// 道理很简单 就是让他们从在相交前从同一个偏移量开始遍历 时间复杂度为O(m+n)
	for {

		if sentinelA == sentinelB {
			return sentinelA
		}

		if sentinelA.Next != nil {
			sentinelA = sentinelA.Next
		} else {
			if changedA == true {
				// 如果已经交换了链表 则表明没有交点
				return nil
			}
			sentinelA = headB
			changedA = true
		}
		if sentinelB.Next != nil {
			sentinelB = sentinelB.Next
		} else {
			if changedB == true {
				// 如果已经交换了链表 则表明没有交点
				return nil
			}
			sentinelB = headA
			changedB = true
		}
	}

}
func main() {

	ListNode4 := ListNode{
		Val:  -4,
		Next: nil,
	}
	ListNode3 := ListNode{
		Val:  0,
		Next: &ListNode4,
	}
	ListNode2 := ListNode{
		Val:  2,
		Next: &ListNode3,
	}
	ListNode1 := ListNode{
		Val:  3,
		Next: &ListNode2,
	}
	ListNode0 := ListNode{
		Val:  3,
		Next: &ListNode3,
	}
	ListNode4.Next = &ListNode2
	getIntersectionNode(&ListNode1, &ListNode0)
	time.Sleep(1)
}
