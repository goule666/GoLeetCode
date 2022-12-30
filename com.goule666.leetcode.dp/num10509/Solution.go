/*
author niewenlong
date 2022/12/30 20:46
description 509. 斐波那契数
leetCodeUrl https://leetcode.cn/problems/fibonacci-number/
*/

package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}
	var dp = make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[len(dp)-1]
}

func main() {
	fmt.Println(fib(3))
}
