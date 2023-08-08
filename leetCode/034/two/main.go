package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	leftleft := &TreeNode{Val: 10}
	leftright := &TreeNode{Val: 10}
	left := &TreeNode{Val: 10, Left: leftleft, Right: leftright}
	right := &TreeNode{Val: 20}

	root := &TreeNode{Val: 0, Left: left, Right: right}

	fmt.Println(pathSum(root, 20))
}
