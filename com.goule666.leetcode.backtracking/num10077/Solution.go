/*
author niewenlong
date 2022/12/21 14:58
description 77. 组合
leetCodeUrl https://leetcode.cn/problems/combinations/
*/

package main

import (
	"fmt"
)

func combine(n int, k int) [][]int {
	var res [][]int
	var temp []int
	var backtracking func(n, k, startIndex int)
	backtracking = func(n, k, startIndex int) {
		if len(temp) == k {
			temp1 := make([]int, k)
			copy(temp1, temp)
			res = append(res, temp1)
			return
		}
		for i := startIndex; i <= n; i++ {
			temp = append(temp, i)
			backtracking(n, k, i+1)
			temp = temp[:len(temp)-1]
		}
	}
	backtracking(n, k, 1)
	return res
}

func main() {
	fmt.Println(combine(4, 2))
}
