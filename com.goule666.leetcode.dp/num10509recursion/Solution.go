/*
author niewenlong
date 2022/12/30 20:46
description 509. 斐波那契数-递归写法
leetCodeUrl https://leetcode.cn/problems/fibonacci-number/
*/

package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(fib(3))
}
