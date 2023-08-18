package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	prt := strings.Builder{}
	for node != nil {
		prt.WriteString(strconv.Itoa(node.Val))
		node = node.Next
	}
	return prt.String()
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	lenList := 0
	now := head
	var last *ListNode
	var next *ListNode
	for now != nil {
		next = now.Next
		now.Next = last
		last = now
		now = next
		lenList++
	}
	return last
}

func main() {
	b := ListNode{Val: 2}
	a := ListNode{Val: 1, Next: &b}
	fmt.Println(reverseList(&a))
}
