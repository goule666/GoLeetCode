/*
author niewenlong
date 2022/12/17 22:23
description 257. 二叉树的所有路径
leetCodeUrl https://leetcode.cn/problems/binary-tree-paths/
*/

package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var res []string
	var rever func(root *TreeNode, s string)
	rever = func(root *TreeNode, s string) {
		if root.Left == nil && root.Right == nil {
			s = s + strconv.Itoa(root.Val)
			res = append(res, s)
			return
		}
		s = s + strconv.Itoa(root.Val) + "->"
		if root.Left != nil {
			rever(root.Left, s)
		}
		if root.Right != nil {
			rever(root.Right, s)
		}
	}
	rever(root, "")
	return res
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(binaryTreePaths(root))
}
