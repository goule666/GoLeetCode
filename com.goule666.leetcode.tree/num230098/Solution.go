/*
author niewenlong
date 2022/12/19 21:31
description 98. 验证二叉搜索树
leetCodeUrl https://leetcode.cn/problems/validate-binary-search-tree/
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

func isValidBST(root *TreeNode) bool {
	var node *TreeNode
	var travel func(root *TreeNode) bool
	travel = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		leftValid := travel(root.Left)
		if node != nil && node.Val >= root.Val {
			return false
		}
		node = root
		rightValid := travel(root.Right)
		return leftValid && rightValid
	}
	return travel(root)
}

func main() {
	root := &TreeNode{Val: 0}
	fmt.Println(isValidBST(root))
}
