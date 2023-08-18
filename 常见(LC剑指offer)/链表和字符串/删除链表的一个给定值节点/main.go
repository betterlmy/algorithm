package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	now := head
	next := now.Next
	for next.Val != val {
		now = next
		next = now.Next
	}
	now.Next = next.Next
	next.Next = nil
	return head

}
