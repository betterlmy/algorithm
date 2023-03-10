package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// checkLeaf 判断是否是叶子结点
func checkLeaf(node *TreeNode) bool {
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

func insert(subStr string, node *TreeNode, res *[]string) {
	if subStr == "" {
		subStr += strconv.Itoa(node.Val)
	} else {
		subStr += "->" + strconv.Itoa(node.Val)
	}

	if checkLeaf(node) {
		*res = append(*res, subStr)
		return
	}
	if node.Left != nil {
		insert(subStr, node.Left, res)
	}

	if node.Right != nil {
		insert(subStr, node.Right, res)
	}

}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if checkLeaf(root) {
		res = append(res, strconv.Itoa(root.Val))
		return res
	}
	insert("", root, &res)
	return res
}

func main() {
	node3 := TreeNode{Val: 5,
		Left:  nil,
		Right: nil,
	}
	node2 := TreeNode{Val: 3,
		Left:  nil,
		Right: nil,
	}
	node1 := TreeNode{Val: 2,
		Left:  nil,
		Right: &node3,
	}
	node0 := TreeNode{Val: 1,
		Left:  &node1,
		Right: &node2,
	}
	_ = node0

	x := binaryTreePaths(&node0)

	for i := range x {
		fmt.Println(x[i])
	}

}
