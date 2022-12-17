/*
author niewenlong
date 2022/12/17 19:58
description 101. 对称二叉树（队列法）
leetCodeUrl https://leetcode.cn/problems/symmetric-tree/
*/

package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	queue := list.New()
	queue.PushFront(root.Left)
	queue.PushBack(root.Right)

	for queue.Len() > 0 {
		left := queue.Remove(queue.Front()).(*TreeNode)
		right := queue.Remove(queue.Back()).(*TreeNode)
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue.PushFront(left.Left)
		queue.PushFront(left.Right)
		queue.PushBack(right.Right)
		queue.PushBack(right.Left)
	}
	return true
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}
	fmt.Println(isSymmetric(root))
}
