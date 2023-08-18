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

// constructBinaryTree 通过数组构建二叉树 不存在为-1
func constructBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	nowIndex := 1
	NodeList := []*TreeNode{root}
	// 每个节点的左右孩子分别是2index+1和2index+2
	for len(NodeList) > 0 && nowIndex < len(nums) {
		node := NodeList[0]
		NodeList = NodeList[1:]
		if nums[nowIndex] >= 0 {
			node.Left = &TreeNode{Val: nums[nowIndex]}
			NodeList = append(NodeList, node.Left)
		}
		nowIndex++
		if nowIndex < len(nums) && nums[nowIndex] >= 0 {
			node.Right = &TreeNode{Val: nums[nowIndex]}
			NodeList = append(NodeList, node.Right)
		}
		nowIndex++
	}
	return root
}

// printTree 层次遍历
func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			fmt.Print("null ")
		} else {
			fmt.Print(strconv.Itoa(node.Val) + " ")
			queue = append(queue, node.Left, node.Right)
		}
	}
	fmt.Println()
}
