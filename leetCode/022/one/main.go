package main

import "time"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
func detectCycle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	sentinel1Node := head

	sentinel2Node := head

	for {
		if sentinel1Node.Next == nil || sentinel2Node.Next == nil || sentinel2Node.Next.Next == nil {
			return nil
		}
		sentinel1Node = sentinel1Node.Next      // 慢指针
		sentinel2Node = sentinel2Node.Next.Next // 快指针

		if sentinel1Node == sentinel2Node {
			//a = c+(n-1)(b+c)
			sentinel3Node := head
			for {
				if sentinel3Node == sentinel1Node {
					return sentinel3Node
				}
				sentinel3Node = sentinel3Node.Next
				sentinel1Node = sentinel1Node.Next
			}

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
	ListNode4.Next = &ListNode2
	detectCycle(&ListNode1)
	time.Sleep(1)
}
