/*
author niewenlong
date 2022/12/30 18:03
description 46. 全排列
leetCodeUrl https://leetcode.cn/problems/permutations/
*/

package main

func permute(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	var res [][]int
	var path []int
	var used = make([]bool, 21)
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		for i := 0; i < len(nums); i++ {
			if used[nums[i]+10] {
				continue
			}
			path = append(path, nums[i])
			used[nums[i]+10] = true
			dfs()
			used[nums[i]+10] = false
			path = path[:len(path)-1]
		}
	}
	dfs()
	return res
}

func main() {
	permute([]int{1, 2, 3})
}
