package main

import (
	"fmt"
	"math"
)

/*
题目:
输入一棵树 T，你需要删除一条边，这棵树会被分成A 和 B 两棵树。
你需要让两部分的**节点数**的差的绝对值| |A|-|B| |尽可能小。
输出最小的| |A|-|B| |和最优方案的数量。

输入描述:
第一行一个整数 n表示节点的数量，节点从1 到 n编号。

接下来n-1行每行两个正整数 s t，表示s的父亲是t。

输入保证是一棵树。

对于所有数据 1<=n<=100000。

输出描述:
输出一行，两个整数，用空格分开，分别表示最优解和最优方案数。
*/

type Node struct {
	id       int   // 当前节点的id
	pid      int   // 父节点id
	size     int   // 子树的大小(包含自己)
	subTrees []int // 子树的id列表
	enable   bool
}

func solution() {
	n := 0 // 表示节点数量
	fmt.Scan(&n)
	tree := make([]Node, n+1)
	tree[1] = Node{id: 1}
	for i := 0; i < n-1; i++ {
		// 扫描n-1行构建整个树
		s, t := 0, 0
		fmt.Scan(&s, &t) //
		// 记录一下 父子节点信息
		tree[s] = Node{
			id:   s,
			pid:  t,
			size: 1,
		}
		tree[t].subTrees = append(tree[t].subTrees, s)
	}
	tree[1].size = update(&tree, 1) // 更新每个节点的size(子树大小)

	// 枚举每条边 假设这条边已经被删掉了
	minDiff := n
	count := 0
	for i := 2; i < len(tree); i++ {
		parent := tree[i].pid
		size1 := tree[i].size // 把i这条边间掉
		size2 := tree[parent].size
		a := size2 - size1
		b := size1 - a
		amb := int(math.Abs(float64(b)))
		be := minDiff
		minDiff = min(amb, minDiff)
		if be == minDiff {
			count++
		} else {
			count = 1
		}

	}

	fmt.Println(minDiff, count)

}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func update(tree *[]Node, nodeNumber int) int {
	// 思路:每个node的size = 所有子树的size+1
	// 递归 因为要求所有子树的size确定 所以设置一个bool变量来判断是否已经被统计过

	node := (*tree)[nodeNumber]
	if node.subTrees == nil {
		// 叶子结点
		node.enable = true
		return 1
	}

	size := 1
	for _, subNode := range node.subTrees {
		if (*tree)[subNode].enable == false {
			// 未被访问过
			tmp := update(tree, subNode)
			(*tree)[subNode].size = tmp
			(*tree)[subNode].enable = true
			size += tmp
		} else {
			size += (*tree)[subNode].size
		}
	}
	return size
}

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*2) // 4s超时
	//defer cancel()
	solution()

	//select {
	//case <-ctx.Done():
	//	return
	//}

}
