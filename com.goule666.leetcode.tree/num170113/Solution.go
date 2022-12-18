/*
author niewenlong
date 2022/12/18 13:25
description 113. 路径总和 II
leetCodeUrl https://leetcode.cn/problems/path-sum-ii/
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var arrays []int
	var loop func(root *TreeNode, targetSum int)

	loop = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}
		targetSum = targetSum - root.Val
		if root.Left != nil || root.Right != nil {
			arrays = append(arrays, root.Val)
		}
		if root.Left == nil && root.Right == nil && targetSum == 0 {
			arrays = append(arrays, root.Val)
			temp := make([]int, len(arrays))
			copy(temp, arrays)
			res = append(res, temp)
			arrays = arrays[:len(arrays)-1]
			return
		}
		if root.Left == nil && root.Right == nil && targetSum != 0 {
			return
		}
		loop(root.Left, targetSum)
		loop(root.Right, targetSum)
		arrays = arrays[:len(arrays)-1]
	}
	loop(root, targetSum)

	return res
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
		Right: &TreeNode{
			Val:  8,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 5},
				Right: nil,
			},
		},
	}
	fmt.Println(pathSum(root, 22))
}
