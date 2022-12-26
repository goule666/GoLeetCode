/*
author niewenlong
date 2022/12/26 19:08
description 93. 复原 IP 地址
leetCodeUrl https://leetcode.cn/problems/restore-ip-addresses/
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	if 0 == len(s) {
		return nil
	}
	var res []string
	var path []string

	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		if len(path) == 4 {
			if len(s) == startIndex {
				res = append(res, strings.Join(path, "."))
			}
			return
		}
		if startIndex >= len(s) {
			return
		}
		for i := startIndex; i < len(s); i++ {
			if isValid(s[startIndex : i+1]) {
				path = append(path, s[startIndex:i+1])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}

func isValid(s string) bool {
	//前导不能为0
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	res, err := strconv.Atoi(s)
	if err != nil || res < 0 || res > 255 {
		return false
	}
	return true
}

func main() {
	addresses := restoreIpAddresses("101023")
	fmt.Println(addresses)
}
