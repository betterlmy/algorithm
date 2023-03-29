package main

import (
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// middleNode 快慢指针获取中间节点
func middleNode(head *ListNode) *ListNode {
	slowP := head
	quickP := head
	for {
		if quickP.Next == nil {
			return slowP
		}
		if quickP.Next.Next == nil {
			return slowP.Next
		}
		quickP = quickP.Next.Next
		slowP = slowP.Next
	}

}

// reverseList 反向链表
func reverseList(l *ListNode) *ListNode {
	var preNode *ListNode //
	curNode := l
	for curNode != nil {
		next := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = next
	}

	return preNode
}

func reorderList(head *ListNode) {

	midNode := middleNode(head)
	l1 := head
	l2 := midNode.Next
	midNode.Next = nil // 修改地址 防止反转过度
	l2 = reverseList(l2)
	mergeList(l1, l2)
}

func mergeList(l1 *ListNode, l2 *ListNode) {
	for l1 != nil && l2 != nil {
		tmp1 := l1.Next
		tmp2 := l2.Next
		l1.Next = l2
		l1 = tmp1

		l2.Next = l1
		l2 = tmp2
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
	reorderList(&ListNode1)
	time.Sleep(1)
}
