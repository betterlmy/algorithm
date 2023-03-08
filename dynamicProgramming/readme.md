# 动态规划

淘宝的“双十一”购物节有各种促销活动，比如“满 200 元减 50 元”。假设你女朋友的购物车中有 n 个（n>100）想买的商品，她希望从里面选几个，在凑够满减条件的前提下，让选出来的商品价格总和最大程度地接近满减条件（200 元），这样就可以极大限度地“薅羊毛”。作为程序员的你，能不能编个代码来帮她搞定呢？

## 动态规划学习路线

动态规划比较适合用来**求解最优问题,例如求最大值最小值**等等.

动态规划比较难学,分为三个步骤:

1. 初识动态规划
2. 动态规划理论
3. 动态规划实战

第一节，我会通过两个非常经典的动态规划问题模型，向你展示我们为什么需要动态规划，以及动态规划解题方法是如何演化出来的。

第二节，我会总结动态规划适合解决的问题的特征，以及动态规划解题思路。除此之外，我还会将贪心、分治、回溯、动态规划这四种算法思想放在一起，对比分析它们各自的特点以及适用的场景。

第三节，我会教你应用第二节讲的动态规划理论知识，实战解决三个非常经典的动态规划问题，加深你对理论的理解。弄懂了这三节中的例子，对于动态规划这个知识点，你就算是入门了。

## 0-1 背包问题

对于一组**不同重量**、**不可分割**的物品，我们需要选择一些装入背包，在满足背包最大重量限制的前提下，背包中物品总重量的最大值是多少呢？

### 回溯解决

在回溯法中,0-1背包问题通过穷举法不断获得所有可能的装法,获得最大值.然而回溯法的时间复杂度是指数级的$$O(2^n)$$.

,我们假设背包的最大承载重量是 9。我们有 5 个不同的物品，每个物品的重量分别是 2，2，4，6，3。回溯法的递归树如下:

![img](readme.assets/42ca6cec4ad034fc3e5c0605fbacecea.jpg)

递归树中的每个节点表示一种状态，我们用（i, cw）来表示。其中，i 表示将要决策第几个物品是否装入背包，cw 表示当前背包中物品的总重量。比如，（2，2）表示我们将要决策第 2 个物品是否装入背包，在决策前，背包中物品的总重量是 2。

图中的黄色部分是重复计算的部分,包括后面的分枝.借助“备忘录”的解决方式，记录已经计算好的 f(i, cw)，当再次计算到重复的 f(i, cw) 的时候，可以直接从备忘录中取出来用，就不用再递归计算了，这样就可以避免冗余计算。

### 动态规划解决

动态规划把整个求解过程分为n个阶段,每个阶段会决策一个物品是否放到背包中.每个物品决策(放或者不放)之后,背包中的物品的重量会有多个不同的情况,也就是多个不同的状态,对应到递归树中,就是有很多不同的节点.

我们把每一层重复的状态(节点)合并,只记录不同的状态,然后基于上一层的状态集合来推导下一层的状态集合.

我们可以通过合并每一层重复的状态，这样就保证每一层不同状态的个数都不会超过 w 个（w 表示背包的承载重量），也就是例子中的 9。于是，我们就成功避免了每层状态个数的指数级增长。

我们用一个二维数组 states\[n\]\[w+1\]，来记录每层可以达到的不同状态。

> n表示n个物品,每次放或者不放,w+1表示总共可能出现的重量情况(最重是w),+1是因为可能全部都不放

第 0 个（下标从 0 开始编号）物品的重量是 2，要么装入背包，要么不装入背包，决策完之后，会对应背包的两种状态，背包中物品的总重量是 0 或者 2。我们用 states\[0\]\[0\]=true 和 states\[0\]\[2\]=true 来表示这两种状态。

第 1 个物品的重量也是 2，基于之前的背包状态，在这个物品决策完之后，不同的状态有 3 个，背包中物品总重量分别是 0=(0+0)，2=(0+2 or 2+0)，4=(2+2)。我们用 states\[1\]\[0\]=true，states\[1\]\[2\]=true，states\[1\]\[4\]=true 来表示这三种状态。

以此类推，直到考察完所有的物品后，整个 states 状态数组就都计算好了。我们只需要在最后一层，找一个值为 true 的最接近 w（这里是 9）的值，就是背包中物品总重量的最大值。

下图中展示了这个过程.图中 0 表示 false，1 表示 true。

![img](readme.assets/aaf51df520ea6b8056f4e62aed81a5b5.jpg)

![img](readme.assets/bbbb934247219db8299bd46dba9dd47e.jpg)

[代码文件](KnapsackProblem/main.go)

```go
package main

import (
	"fmt"
)

var (
	capacity     = 9
	weights      = []int{2, 2, 4, 6, 3}
	numItems     = len(weights)
	stateChart   = make([][]int, numItems)
	knownWeights = make([]int, 0)
)

func dynamicProgramming() {
	for i, weight := range weights {
		if i == 0 {
			stateChart[i][0] = 1
			knownWeights = append(knownWeights, 0)
			if weight <= capacity {
				stateChart[i][weight] = 1
				knownWeights = append(knownWeights, weight)
			}
		} else {
			for _, knownWeight := range knownWeights {
				stateChart[i][knownWeight] = 1
				newWeight := knownWeight + weight
				if newWeight <= capacity && !contain(newWeight, knownWeights) {
					knownWeights = append(knownWeights, newWeight)
					stateChart[i][newWeight] = 1
				}
			}
		}

	}

}

func contain(weight int, ints []int) bool {
	for _, i2 := range ints {
		if weight == i2 {
			return true
		}
	}
	return false
}

func main() {
	for i := range stateChart {
		stateChart[i] = make([]int, capacity+1)
	}
	dynamicProgramming()
	for i, rows := range stateChart {
		if i == 0 {
			fmt.Println("   0 1 2 3 4 5 6 7 8 9 ")
		}
		fmt.Printf("%d  ", i)
		for _, data := range rows {
			fmt.Print(data, " ")
		}
		fmt.Println()
	}
	x := stateChart[len(stateChart)-1]
	for i := len(x) - 1; i >= 0; i-- {
		if x[i] == 1 {
			fmt.Println("最多的容量是:", i)
			break
		}
	}

}

```

如果在0-1背包问题中不仅仅有数量和重量两个条件,再加入一个价值的条件,要求如何放在满足重量的条件下,尽可能获得较大的价值.

使用DP算法只需要将stateChart中的元素从0和1改为价值,每次获取一个最高的价值并更新,将默认值从0修改为-1即可

大部分动态规划能解决的问题，都可以通过回溯算法来解决，只不过回溯算法解决起来效率比较低，时间复杂度是指数级的。动态规划算法，在执行效率方面，要高很多。尽管执行效率提高了，但是动态规划的空间复杂度也提高了，所以，很多时候，我们会说，动态规划是一种**空间换时间**的算法思想。

