/*
author niewenlong
date 2022/12/20 15:04
description 538. 把二叉搜索树转换为累加树
leetCodeUrl https://leetcode.cn/problems/convert-bst-to-greater-tree/
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	count := 0
	var traversal func(cur *TreeNode)
	traversal = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		traversal(cur.Right)
		cur.Val += count
		count = cur.Val
		traversal(cur.Left)
	}
	traversal(root)
	return root
}

func main() {
	convertBST(nil)
}
