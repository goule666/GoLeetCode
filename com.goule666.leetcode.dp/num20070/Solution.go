/*
author niewenlong
date 2022/12/30 21:09
description 70. 爬楼梯
leetCodeUrl https://leetcode.cn/problems/climbing-stairs/
*/

package main

import "fmt"

func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	var dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	fmt.Println(climbStairs(2))
}
