/*
author niewenlong
date 2022/12/19 22:39
description 530. 二叉搜索树的最小绝对差
leetCodeUrl https://leetcode.cn/problems/minimum-absolute-difference-in-bst/
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	var result int
	var prev *TreeNode
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		if prev != nil {
			if result == 0 {
				result = root.Val - prev.Val
			} else {
				result = min(result, root.Val-prev.Val)
			}
		}
		prev = root
		traversal(root.Right)
	}
	traversal(root)
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 6},
	}
	fmt.Println(getMinimumDifference(root))
}
