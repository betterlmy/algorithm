package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() (_ Codec) {
	return
}

func (this *Codec) serialize(root *TreeNode) string {
	// 层次遍历
	series := ""
	nodeList := make([]*TreeNode, 0)
	nodeList = append(nodeList, root)

	for len(nodeList) > 0 {
		node := nodeList[0]
		if node != nil {
			series += strconv.Itoa(node.Val) + ","
			nodeList = append(nodeList, node.Left)
			nodeList = append(nodeList, node.Right)
		} else {
			series += "nil,"
		}
		nodeList = nodeList[1:]

	}

	return series
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	datas := strings.Split(data, ",")
	datas = datas[:len(datas)-1]
	if datas[0] == "nil" {
		return nil
	}

	root := &TreeNode{}
	nodeList := make([]*TreeNode, 0)
	nodeList = append(nodeList, root)
	lastNode := &TreeNode{}
	left := true // 当前正在处理left?
	for _, v := range datas {
		left = !left
		if v != "nil" {
			x, _ := strconv.Atoi(v)
			nodeList[0].Val = x
			nodeList[0].Left = &TreeNode{}
			nodeList[0].Right = &TreeNode{}
			nodeList = append(nodeList, nodeList[0].Left)
			nodeList = append(nodeList, nodeList[0].Right)
			if !left {
				lastNode = nodeList[0]
			}

			nodeList = nodeList[1:]
		} else {
			if left {
				lastNode.Left = nil
			} else {
				lastNode.Right = nil
			}
			nodeList = nodeList[1:]
		}
	}

	return root
}

func main() {
	codec := Codec{}
	node1 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	node2 := &TreeNode{
		Val:   2,
		Left:  node1,
		Right: nil,
	}
	node3 := &TreeNode{
		Val:   3,
		Left:  node2,
		Right: nil,
	}
	series := codec.serialize(node3)
	fmt.Println(series)
	codec.deserialize(series)
}
