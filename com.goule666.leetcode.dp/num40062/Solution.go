/*
author niewenlong
date 2022/12/31 15:20
description 62. 不同路径
leetCodeUrl https://leetcode.cn/problems/unique-paths/
*/

package main

import "fmt"

func uniquePaths(m int, n int) int {
	var dp = make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 7))
}
