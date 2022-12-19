/*
author niewenlong
date 2022/12/19 23:11
description 501. 二叉搜索树中的众数
leetCodeUrl https://leetcode.cn/problems/find-mode-in-binary-search-tree/
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var count int
	var maxCount int
	var prev *TreeNode
	var traversal func(root *TreeNode)
	traversal = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		traversal(cur.Left)
		if prev == nil {
			count = 1
		} else if prev.Val == cur.Val {
			count++
		} else {
			count = 1
		}
		if count == maxCount {
			res = append(res, cur.Val)
		}
		if count > maxCount {
			maxCount = count
			res = []int{cur.Val}
		}
		prev = cur
		traversal(cur.Right)
	}
	traversal(root)
	return res
}

func main() {
	findMode(nil)
}
