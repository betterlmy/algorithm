package main

import "fmt"

/*
好串定义为"当且仅当该字符串不包含长度超过1的回文子串"
给一个仅有小写字母组成的字符串,问这个字符串中有多少个子串是好串
*/

func main() {
	n := 0
	fmt.Scan(&n)
	// 捕获字符串
	var str string
	fmt.Scan(&str)
	fmt.Println(goodSubstr(str))

}
func goodSubstr(s string) int {

	// 动态规划
	// dp[i] 表示以第i个字符结尾的最长非回文子串的长度 如果一个子串的长度等于dp[i] 那么他就是好串
	n := len(s)
	dp := make([]int, n) //第i个字符结尾的最长非回文子串的长度
	res := 0
	for i := 0; i < n; i++ {
		// init
		dp[i] = 1 // 默认就是自己 长度为1

		if i > 0 && s[i] == s[i-1] {
			// 连续相同的情况 aa dp[i]=dp[i-1]
			dp[i] = dp[i-1]
		} else if i > 1 && s[i] == s[i-2] {
			// aaba
			dp[i] = dp[i-2] + 2
		} else if i > 1 {
			if s[i-1] == s[i+1] {
				dp[i] = 2
			}
		}

		if dp[i] == i+1 || !isHuiwen(s[i-dp[i]+1:i+1]) {
			res++
		}

	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isHuiwen(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
