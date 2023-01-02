/*
author niewenlong
date 2023/1/2 13:26
description 343. 整数拆分
leetCodeUrl https://leetcode.cn/problems/integer-break/
*/

package main

import "fmt"

func integerBreak(n int) int {
	var dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	//1 j i n
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(max(dp[i-j]*j, (i-j)*j), dp[i])
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(integerBreak(3))
}
