package main

import (
	"fmt"
)

type Node struct {
	id       int   // 当前节点的id
	pid      int   // 父节点id
	size     int   // 子树的末尾几个0
	subTrees []int // 子树的id列表
	enable   bool
}

func solution() {
	n := 0 // 表示节点数量
	fmt.Scan(&n)

	tree := make([]Node, n+1)
	tree[1] = Node{id: 1}
	for i := 0; i < n-1; i++ {

	}

}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {

	solution()

}
