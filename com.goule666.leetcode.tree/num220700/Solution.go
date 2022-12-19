/*
author niewenlong
date 2022/12/19 21:26
description 700. 二叉搜索树中的搜索
leetCodeUrl https://leetcode.cn/problems/search-in-a-binary-search-tree/
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return nil
}

func main() {
	searchBST(nil, 0)
}
