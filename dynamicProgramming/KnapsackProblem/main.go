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
