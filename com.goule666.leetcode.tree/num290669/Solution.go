/*
author niewenlong
date 2022/12/20 14:18
description 669. 修剪二叉搜索树
leetCodeUrl https://leetcode.cn/problems/trim-a-binary-search-tree/
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val < low {
		node := trimBST(root.Right, low, high)
		return node
	}
	if root.Val > high {
		node := trimBST(root.Left, low, high)
		return node
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

func main() {
	trimBST(nil, 0, 0)
}
