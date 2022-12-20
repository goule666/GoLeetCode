/*
author niewenlong
date 2022/12/19 23:45
description 235. 二叉搜索树的最近公共祖先（迭代法）
leetCodeUrl https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	l, f := sortNode(p, q)
	for root != nil {
		if root.Val > f.Val {
			root = root.Left
		} else if root.Val < l.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}

func sortNode(p, q *TreeNode) (l, f *TreeNode) {
	if p.Val < q.Val {
		return p, q
	} else {
		return q, p
	}
}

func main() {
	p := &TreeNode{Val: 3}
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 1},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   8,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 9},
			},
		},
	}
	fmt.Println(lowestCommonAncestor(root, p, root.Left))
}
