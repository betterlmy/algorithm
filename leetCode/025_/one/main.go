package main

import (
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// reverseList
	l1 = reverseList(l1)
	l2 = reverseList(l2)

	var resList = new(ListNode)
	head := ListNode{
		Val:  0,
		Next: resList,
	}
	c := 0 // 进位
	finishedList := 0
	for {
		sum := l1.Val + l2.Val + c
		if sum >= 10 {
			c = 1
			sum -= 10
		} else {
			c = 0
		}
		resList.Val = sum

		l1 = l1.Next
		l2 = l2.Next
		if l1 == nil {
			finishedList += 1
		}
		if l2 == nil {
			finishedList += 2
		}
		if finishedList > 0 {
			break
		}
		resList.Next = new(ListNode)
		resList = resList.Next
	}

	switch finishedList {
	case 1:
		// l2还有
		for {
			resList.Next = new(ListNode)
			resList = resList.Next
			sum := l2.Val + c
			if sum >= 10 {
				c = 1
				sum -= 10
			} else {
				c = 0
			}
			resList.Val = sum
			l2 = l2.Next
			if l2 == nil {
				break
			}
		}

	case 2:
		// l1还有
		for {
			resList.Next = new(ListNode)
			resList = resList.Next
			sum := l1.Val + c
			if sum >= 10 {
				c = 1
				sum -= 10
			} else {
				c = 0
			}
			resList.Val = sum
			l1 = l1.Next
			if l1 == nil {
				break
			}
		}
	}
	if c == 1 {
		resList.Next = new(ListNode)
		resList = resList.Next
		resList.Val = 1
	}

	return reverseList(head.Next)
}

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

func main() {

	ListNode1 := ListNode{
		Val:  1,
		Next: nil,
	}

	ListNode5 := ListNode{
		Val:  9,
		Next: nil,
	}
	ListNode6 := ListNode{
		Val:  9,
		Next: &ListNode5,
	}

	x := addTwoNumbers(&ListNode1, &ListNode6)
	_ = x
	time.Sleep(1)
}
