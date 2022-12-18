/*
author niewenlong
date 2022/12/18 12:55
description 404. 左叶子之和
leetCodeUrl https://leetcode.cn/problems/sum-of-left-leaves/
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftValue := sumOfLeftLeaves(root.Left)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftValue = root.Left.Val
	}

	rightValue := sumOfLeftLeaves(root.Right)
	return leftValue + rightValue
}

func main() {
	sumOfLeftLeaves(nil)
}
