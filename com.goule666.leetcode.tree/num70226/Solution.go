/*
author niewenlong
date 2022/12/16 16:53
description 226. 翻转二叉树
leetCodeUrl https://leetcode.cn/problems/invert-binary-tree/
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var loop func(root *TreeNode)
	loop = func(root *TreeNode) {
		if root == nil {
			return
		}
		root.Left, root.Right = root.Right, root.Left
		loop(root.Left)
		loop(root.Right)
	}
	loop(root)

	return root
}

func main() {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Print(invertTree(root))

}
