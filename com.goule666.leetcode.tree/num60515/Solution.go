/*
author niewenlong
date 2022/12/16 15:40
description 515. 在每个树行中找最大值
leetCodeUrl https://leetcode.cn/problems/find-largest-value-in-each-tree-row/
*/

package main

import (
	"container/list"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var maxNum = math.MinInt
		queueSize := queue.Len()
		for i := 0; i < queueSize; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			maxNum = Max(maxNum, node.Val)
		}
		res = append(res, maxNum)
	}
	return res
}

func Max(values ...int) int {
	m := math.MinInt
	for _, v := range values {
		if v > m {
			m = v
		}
	}
	return m
}

func main() {
	largestValues(nil)
}
