/*
author niewenlong
date 2022/12/17 21:24
description 111.二叉树的最小深度
leetCodeUrl https://leetcode.cn/problems/minimum-depth-of-binary-tree/
*/

package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)

	minDepth := 0
	for queue.Len() > 0 {
		minDepth++
		queueSize := queue.Len()
		for i := 0; i < queueSize; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return minDepth
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return minDepth
}

func main() {
	minDepth(nil)
}
