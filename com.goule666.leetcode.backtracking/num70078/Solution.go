/*
author niewenlong
date 2022/12/26 20:01
description 78. 子集
leetCodeUrl https://leetcode.cn/problems/subsets/
*/

package main

import "fmt"

func subsets(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	var res [][]int
	var path []int
	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		if path == nil {
			res = append(res, []int{})
		} else {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}

		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1) //1//2//3
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
