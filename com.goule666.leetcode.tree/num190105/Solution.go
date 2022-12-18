/*
author niewenlong
date 2022/12/18 15:17
description 105. 从前序与中序遍历序列构造二叉树
leetCodeUrl https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 || inorder == nil || len(inorder) == 0 {
		return nil
	}
	//定root元素
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}

	//根据root元素切割中序集合
	sliceCutInOrderIndex := 0
	for index, ele := range inorder {
		if ele == root.Val {
			sliceCutInOrderIndex = index
			break
		}
	}
	leftInorder, rightInorder := inorder[0:sliceCutInOrderIndex], inorder[sliceCutInOrderIndex+1:]
	leftPreorder, rightPreorder := preorder[1:len(leftInorder)+1], preorder[len(leftInorder)+1:]

	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)
	return root
}

func main() {
	fmt.Println(buildTree([]int{1, 2}, []int{2, 1}))
}
