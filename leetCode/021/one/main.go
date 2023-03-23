package main

import "time"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentinelNode := ListNode{0, head}
	nodeAddr := make([]*ListNode, 0)
	nowNode := sentinelNode.Next
	lenList := 1
	for {
		nodeAddr = append(nodeAddr, nowNode)
		if nowNode.Next != nil {
			nowNode = nowNode.Next
			lenList++
		} else {
			break
		}
	}
	delPosition := lenList - n
	if delPosition == 0 {
		// 特殊情况
		if len(nodeAddr) == 1 {
			sentinelNode.Next = nil
		} else {
			sentinelNode.Next = nodeAddr[1]
		}
	} else {
		if delPosition+1 >= lenList {
			nodeAddr[delPosition-1].Next = nil
		} else {
			nodeAddr[delPosition-1].Next = nodeAddr[delPosition+1]
		}
		nodeAddr[delPosition].Next = nil
	}

	return sentinelNode.Next
}

func main() {

	ListNode1 := ListNode{
		Val:  1,
		Next: nil,
	}
	removeNthFromEnd(&ListNode1, 1)
	time.Sleep(1)
}
