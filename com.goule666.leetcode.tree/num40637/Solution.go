/*
author niewenlong
date 2022/12/16 14:12
description 637.二叉树的层平均值
leetCodeUrl https://leetcode.cn/problems/average-of-levels-in-binary-tree/
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

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	result := make([]float64, 0)
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		queueSize := queue.Len()
		var sameDeepList []int
		for i := 0; i < queueSize; i++ {
			element := queue.Remove(queue.Front()).(*TreeNode)
			sameDeepList = append(sameDeepList, element.Val)
			if element.Left != nil {
				queue.PushBack(element.Left)
			}
			if element.Right != nil {
				queue.PushBack(element.Right)
			}
		}
		var sameDeepSum int
		for _, ele := range sameDeepList {
			sameDeepSum += ele
		}
		result = append(result, float64(sameDeepSum)/float64(len(sameDeepList)))
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
	fmt.Print(averageOfLevels(root))
}
