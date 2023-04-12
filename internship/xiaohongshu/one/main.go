package main

import (
	"context"
	"fmt"
	"math"
	"time"
)

/*
题目:
数列的定义如下： 数列的第一项为n，以后各项为前一项的平方根，求数列的前m项的和。

输入描述:
输入数据有多组，每组占一行，由两个整数n（n<10000）和m(m<1000)组成，n和m的含义如前所述。

输出描述:
对于每组输入数据，输出该数列的和，每个测试实例占一行，要求精度保留2位小数。


*/

func solution() {
	for {
		n, m := 0, 0
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			break
		}
		sum := float64(n)
		q := float64(n)
		for i := 1; i < m; i++ {
			q = math.Sqrt(q)
			sum += q
		}
		fmt.Printf("%.2f\n", sum)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2) // 4s超时
	defer cancel()
	go solution()

	select {
	case <-ctx.Done():
		return
	}

}
