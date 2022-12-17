/*
author niewenlong
date 2022/12/17 21:18
description 104. 二叉树的最大深度
leetCodeUrl https://leetcode.cn/problems/maximum-depth-of-binary-tree/
*/

package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)

	maxDepth := 0
	for queue.Len() > 0 {
		maxDepth++
		queueSize := queue.Len()
		for i := 0; i < queueSize; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return maxDepth
}

func main() {
	maxDepth(nil)
}
