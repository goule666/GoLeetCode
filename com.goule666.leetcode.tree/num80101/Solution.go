/*
author niewenlong
date 2022/12/17 19:58
description 101. 对称二叉树（迭代法）
leetCodeUrl https://leetcode.cn/problems/symmetric-tree/
*/

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	var loop func(left *TreeNode, right *TreeNode) bool
	loop = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right != nil {
			return false
		} else if left != nil && right == nil {
			return false
		} else if left == nil && right == nil {
			return true
		} else if left.Val != right.Val {
			return false
		}
		outside := loop(left.Left, right.Right)
		inside := loop(left.Right, right.Left)
		return outside && inside
	}
	return loop(root.Left, root.Right)
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}
	fmt.Println(isSymmetric(root))
}
