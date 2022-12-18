/*
author niewenlong
date 2022/12/18 16:05
description 654. 最大二叉树
leetCodeUrl https://leetcode.cn/problems/maximum-binary-tree/
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	//find max val
	maxIndex := 0
	maxVal := 0
	for i, num := range nums {
		if num > maxVal {
			maxIndex = i
			maxVal = num
		}
	}
	//构造root
	root := &TreeNode{Val: maxVal}

	//切割
	leftNums, rightNums := nums[0:maxIndex], nums[maxIndex+1:]

	root.Left = constructMaximumBinaryTree(leftNums)
	root.Right = constructMaximumBinaryTree(rightNums)
	return root
}

func main() {
	fmt.Println(constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}))
}
