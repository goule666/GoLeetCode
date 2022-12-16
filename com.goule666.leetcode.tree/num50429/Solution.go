/*
author niewenlong
date 2022/12/16 14:48
description 429. N 叉树的层序遍历
leetCodeUrl https://leetcode.cn/problems/n-ary-tree-level-order-traversal/
*/

package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		queueSize := queue.Len()
		var sameDeepList []int
		for i := 0; i < queueSize; i++ {
			elementNode := queue.Remove(queue.Front()).(*Node)
			sameDeepList = append(sameDeepList, elementNode.Val)
			if elementNode.Children != nil {
				for _, child := range elementNode.Children {
					queue.PushBack(child)
				}
			}
		}
		res = append(res, sameDeepList)
	}
	return res
}

func main() {

	root := &Node{
		Val: 1,
		Children: []*Node{{
			Val:      3,
			Children: nil,
		}, {
			Val:      2,
			Children: nil,
		}, {
			Val:      4,
			Children: nil,
		}},
	}
	fmt.Print(levelOrder(root))
}
