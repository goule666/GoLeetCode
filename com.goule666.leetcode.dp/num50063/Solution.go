/*
author niewenlong
date 2022/12/31 15:47
description 63. 不同路径 II
leetCodeUrl https://leetcode.cn/problems/unique-paths-ii/
*/

package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid == nil {
		return 0
	}
	//定义数组
	var dp = make([][]int, len(obstacleGrid))
	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	//初始化
	for i := 0; i < len(obstacleGrid) && obstacleGrid[i][0] != 1; i++ {
		dp[i][0] = 1
	}
	//初始化
	for i := 0; i < len(obstacleGrid[0]) && obstacleGrid[0][i] != 1; i++ {
		dp[0][i] = 1
	}
	for m := 1; m < len(obstacleGrid); m++ {
		for n := 1; n < len(obstacleGrid[0]); n++ {
			if obstacleGrid[m][n] == 1 {
				continue
			}
			dp[m][n] = dp[m-1][n] + dp[m][n-1]
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{1}, {0}}))
}
