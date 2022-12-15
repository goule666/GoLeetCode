/*
author niewenlong
date 2022/12/15 22:15
description 144. 二叉树的前序遍历
leetCodeUrl https://leetcode.cn/problems/binary-tree-preorder-traversal/
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var root = &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(inorderTraversal(root))
}
func inorderTraversal(root *TreeNode) (res []int) {
	var loop func(root *TreeNode)
	loop = func(root *TreeNode) {
		if root == nil {
			return
		}
		loop(root.Left)
		res = append(res, root.Val)
		loop(root.Right)
	}
	loop(root)
	return res
}
