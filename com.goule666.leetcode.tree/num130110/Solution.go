/*
author niewenlong
date 2022/12/17 21:50
description 110. 平衡二叉树
leetCodeUrl https://leetcode.cn/problems/balanced-binary-tree/
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var getHeight func(root *TreeNode) int
	getHeight = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := getHeight(root.Left)
		if l == -1 {
			return -1
		}
		r := getHeight(root.Right)
		if r == -1 {
			return -1
		}
		if l-r > 1 || r-l > 1 {
			return -1
		}
		return max(l, r) + 1
	}

	if getHeight(root) == -1 {
		return false
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	isBalanced(nil)
}
