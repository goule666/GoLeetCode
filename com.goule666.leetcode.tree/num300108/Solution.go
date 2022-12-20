/*
author niewenlong
date 2022/12/20 14:47
description 108. 将有序数组转换为二叉搜索树
leetCodeUrl https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	//find center node
	centerNodeIndex := len(nums) / 2
	root := &TreeNode{Val: nums[centerNodeIndex]}

	root.Left = sortedArrayToBST(nums[0:centerNodeIndex])
	root.Right = sortedArrayToBST(nums[centerNodeIndex+1:])
	return root
}

func main() {
	sortedArrayToBST([]int{-10, -3, 0, 5, 9})
}
