/*
author niewenlong
date 2022/12/30 18:03
description 47. 全排列 II
leetCodeUrl https://leetcode.cn/problems/permutations-ii/
*/

package main

import "sort"

func permuteUnique(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	var path []int
	var used = make([]bool, 10)
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		for i := 0; i < len(nums); i++ {
			if (i > 0 && (nums[i] == nums[i-1]) && !used[i-1]) || used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			dfs()
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs()
	return res
}

func main() {
	permuteUnique([]int{1, 1, 2})
}
