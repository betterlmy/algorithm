package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 新建节点 l1 和l2一起遍历 谁小就加入谁然后让指针后移 如果一个先结束了 另一个直接赋值后续即可
	res := &ListNode{}
	resHead := res
	endList := 1 // 标记谁先结束

	// 特殊情况
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			res.Val = l1.Val
			l1 = l1.Next
			endList = 1
		} else {
			res.Val = l2.Val
			l2 = l2.Next
			endList = 2
		}
		res.Next = &ListNode{}
		res = res.Next
	}
	if endList == 2 {
		res.Val = l1.Val
		res.Next = l1.Next
	} else {
		res.Val = l2.Val
		res.Next = l2.Next
	}
	return resHead
}
