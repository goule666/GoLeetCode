/*
author niewenlong
date 2022/12/17 21:27
description 222. 完全二叉树的节点个数
leetCodeUrl https://leetcode.cn/problems/count-complete-tree-nodes/
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := root.Left
	right := root.Right

	//得到左子树最大长度
	maxLeftTreeHeight := 0
	for left != nil {
		left = left.Left
		maxLeftTreeHeight++
	}

	//得到右子树最大长度
	maxRightTreeHeight := 0
	for right != nil {
		right = right.Right
		maxRightTreeHeight++
	}
	if maxLeftTreeHeight == maxRightTreeHeight {
		return 2<<maxRightTreeHeight - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func main() {
	countNodes(nil)
}
