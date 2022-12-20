/*
author niewenlong
date 2022/12/20 11:14
description 701. 二叉搜索树中的插入操作
leetCodeUrl https://leetcode.cn/problems/insert-into-a-binary-search-tree/
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	var traversal func(cur *TreeNode, val int) *TreeNode

	traversal = func(cur *TreeNode, val int) *TreeNode {
		if cur == nil {
			return &TreeNode{Val: val}
		}
		if cur.Val > val {
			cur.Left = traversal(cur.Left, val)
		}
		if cur.Val < val {
			cur.Right = traversal(cur.Right, val)
		}
		return cur
	}
	return traversal(root, val)
}

func main() {
	insertIntoBST(nil, 1)
}
