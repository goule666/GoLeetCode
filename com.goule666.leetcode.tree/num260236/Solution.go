/*
author niewenlong
date 2022/12/19 23:45
description 236. 二叉树的最近公共祖先
leetCodeUrl https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	leftNode := lowestCommonAncestor(root.Left, p, q)
	rightNode := lowestCommonAncestor(root.Right, p, q)
	if leftNode != nil && rightNode != nil {
		return root
	}
	if leftNode == nil && rightNode != nil {
		return rightNode
	}
	if leftNode != nil && rightNode == nil {
		return leftNode
	}
	return nil
}

func main() {
	p := &TreeNode{Val: 3}
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val:   4,
			Left:  p,
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: &TreeNode{Val: 8},
		},
	}
	fmt.Println(lowestCommonAncestor(root, p, root.Left))
}
