/*
author niewenlong
date 2022/12/21 16:26
description 216. 组合总和 III
leetCodeUrl https://leetcode.cn/problems/combination-sum-iii/
*/

package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var path []int
	var dfs func(k, n, startIndex int)
	dfs = func(k, n, startIndex int) {
		if len(path) == k {
			//判断是否满足条件
			if getSum(path) == n {
				temp := make([]int, k)
				copy(temp, path)
				res = append(res, temp)
			}
			return
		}
		for i := startIndex; i <= 9; i++ {
			path = append(path, i)
			dfs(k, n, i+1) //1 //3
			path = path[:len(path)-1]
		}
	}
	dfs(k, n, 1)
	return res
}

func getSum(path []int) int {
	sum := 0
	for _, i2 := range path {
		sum += i2
	}
	return sum
}

func main() {
	fmt.Println(combinationSum3(2, 4))
}
