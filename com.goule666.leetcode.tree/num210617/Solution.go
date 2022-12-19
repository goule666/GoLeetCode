/*
author niewenlong
date 2022/12/18 16:05
description 617.合并二叉树
leetCodeUrl https://leetcode.cn/problems/merge-two-binary-trees/
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val = root1.Val + root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

func main() {
	fmt.Println(mergeTrees(nil, nil))
}
