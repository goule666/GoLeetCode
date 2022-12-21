/*
author niewenlong
date 2022/12/21 16:51
description 17. 电话号码的字母组合
leetCodeUrl https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
*/

package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var res []string
	var path string
	var data = make(map[string][]string)
	data["2"] = []string{"a", "b", "c"}
	data["3"] = []string{"d", "e", "f"}
	data["4"] = []string{"g", "h", "i"}
	data["5"] = []string{"j", "k", "l"}
	data["6"] = []string{"m", "n", "o"}
	data["7"] = []string{"p", "q", "r", "s"}
	data["8"] = []string{"t", "u", "v"}
	data["9"] = []string{"w", "x", "y", "z"}

	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		if len(path) == len(digits) {
			res = append(res, path)
			return
		}
		var item string
		for i, elem := range digits {
			if i == startIndex {
				item = string(elem)
				break
			}
		}
		for _, s := range data[item] {
			path = path + s
			dfs(startIndex + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

func main() {
	fmt.Println(letterCombinations("23"))
}
