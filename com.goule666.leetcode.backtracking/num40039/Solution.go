/*
author niewenlong
date 2022/12/21 17:59
description 39. 组合总和
leetCodeUrl https://leetcode.cn/problems/combination-sum/
*/

package main

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	var res [][]int
	var path []int
	var sum int
	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		if sum > target {
			return
		} else if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := startIndex; i < len(candidates); i++ {
			sum += candidates[i]
			path = append(path, candidates[i])
			dfs(i)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

func main() {
	combinationSum([]int{2, 3, 6, 7}, 7)
}
