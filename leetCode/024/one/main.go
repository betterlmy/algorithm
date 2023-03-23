package main

import "time"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var prevNode *ListNode
	curNode := head

	for curNode != nil {
		next := curNode.Next    // 保存下一个要遍历的节点
		curNode.Next = prevNode // 将当前节点指向上一个节点
		prevNode = curNode      // 上一个节点往后移
		curNode = next          // 当前的节点也往后移动
	}
	return prevNode
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

	x := reverseList(&ListNode1)
	_ = x
	time.Sleep(1)
}
