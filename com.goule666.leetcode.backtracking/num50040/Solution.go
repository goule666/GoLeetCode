/*
author niewenlong
date 2022/12/21 19:49
description 40. 组合总和 II
leetCodeUrl https://leetcode.cn/problems/combination-sum-ii/
*/

package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var path []int
	var sum int
	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		if sum > target {
			return
		}
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := startIndex; i < len(candidates); i++ {
			if i > startIndex && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			dfs(i + 1)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}

	}
	dfs(0)
	return res
}

func main() {
	combinationSum2([]int{1, 1, 2}, 3)
}
