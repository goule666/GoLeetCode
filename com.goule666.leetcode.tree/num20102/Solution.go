/*
author niewenlong
date 2022/12/15 23:24
description 102. 二叉树的层序遍历
leetCodeUrl https://leetcode.cn/problems/binary-tree-level-order-traversal/
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

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := list.New()
	queue.PushBack(root)
	height := 0

	for queue.Len() != 0 {
		queueSize := queue.Len()
		var arrInts []int
		for i := 0; i < queueSize; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			arrInts = append(arrInts, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result = append(result, arrInts)
		height++
	}
	return result
}

func main() {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Print(levelOrder(root))
}
