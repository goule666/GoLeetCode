/*
author niewenlong
date 2022/12/30 21:24
description 746. 使用最小花费爬楼梯
leetCodeUrl https://leetcode.cn/problems/min-cost-climbing-stairs/
*/

package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	var dp = make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}
