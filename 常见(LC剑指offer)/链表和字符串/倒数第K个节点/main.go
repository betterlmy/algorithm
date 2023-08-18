package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	// 从前往后 用一个res指针指向当前节点的前面第k个节点 不断遍历 直到指向nil
	var res = head
	for i := 0; i < k; i++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	for head != nil {
		head = head.Next
		res = res.Next
	}

	return res
}
