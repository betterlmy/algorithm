package main

import (
	"fmt"
)

//给定一个字符串数组words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。
//链接：https://leetcode.cn/problems/aseY1I
/*
1. 子串
2. 计算长度的乘积
3. 不包含相同字符

思路:
1. 暴力法 O(n^2) 不写了
2. 位运算?
由于每个单词只包含小写字母,共26个.因此可以用位掩码的低26位来表示每个字母是否在这个单词中出现.
One-hot?


*/

func maxProduct(words []string) int {
	//
	res := 0
	masks := make([]int, len(words))
	for i, word := range words {
		// 将每个单词都存在一个one-hot向量中 masks存着所有的one-hot向量
		for _, char := range word {
			// char = 'a'时 左移0位,
			charBit := 1 << (char - 'a')
			masks[i] = masks[i] | charBit // 通过或运算将10或者01替换为1(只存在00 01 10 不存在11)
		}
	}

	// Check
	//for i := 25; i >= 0; i-- {
	//	fmt.Print(string('a' + i))
	//}
	//fmt.Println()
	//binaryStr := fmt.Sprintf("%b", masks[0])
	//paddedBinary := fmt.Sprintf("%0*s", 26, binaryStr)
	//fmt.Println(paddedBinary)
	//fmt.Println(words[0])
	length := len(masks)
	for i, mask := range masks {
		for j := i + 1; j < length; j++ {
			if mask&masks[j] == 0 {
				// 两者没有交叉

				nowRes := len(words[i]) * len(words[j])

				if res < nowRes {
					res = nowRes
				}
			}
		}
	}
	return res
}

func main() {
	x := []string{
		"asdfff",
		"g",
	}
	fmt.Println(maxProduct(x))
}
