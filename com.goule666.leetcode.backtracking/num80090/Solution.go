/*
author niewenlong
date 2022/12/26 20:49
description 90.子集II
leetCodeUrl https://leetcode.cn/problems/subsets-ii/
*/

package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	sort.Ints(nums)
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
			if i > startIndex && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
