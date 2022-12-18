/*
author niewenlong
date 2022/12/18 13:25
description 112. 路径总和
leetCodeUrl https://leetcode.cn/problems/path-sum/
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum = targetSum - root.Val
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}
	if root.Left == nil && root.Right == nil && targetSum != 0 {
		return false
	}

	if hasPathSum(root.Left, targetSum) {
		return true
	}
	if hasPathSum(root.Right, targetSum) {
		return true
	}
	return false
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
			Right: nil,
		},
		Right: nil,
	}
	fmt.Println(hasPathSum(root, 22))
}
