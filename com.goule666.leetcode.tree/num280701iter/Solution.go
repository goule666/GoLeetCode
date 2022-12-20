/*
author niewenlong
date 2022/12/20 11:14
description 701. 二叉搜索树中的插入操作（迭代法）
leetCodeUrl https://leetcode.cn/problems/insert-into-a-binary-search-tree/
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	cur := root
	for cur != nil {
		if cur.Val > val {
			if cur.Left == nil {
				break
			}
			cur = cur.Left
		}
		if cur.Val < val {
			if cur.Right == nil {
				break
			}
			cur = cur.Right
		}
	}
	node := &TreeNode{Val: val}
	if cur.Val < val {
		cur.Right = node
	} else {
		cur.Left = node
	}
	return root
}

func main() {
	insertIntoBST(nil, 1)
}
