package main

// 快慢指针
type ListNode struct {
	Val  int
	Next *ListNode
}

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

func main() {

}
