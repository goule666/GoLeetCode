/*
author niewenlong
date 2022/12/18 15:17
description 106. 从中序与后序遍历序列构造二叉树
leetCodeUrl https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if inorder == nil || postorder == nil {
		return nil
	}
	//从后续遍历集合中得到根节点
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	if len(postorder) == 1 {
		return root
	}
	//根据根节点 切割中序遍历集合
	sliceCutIndex := 0
	for index, ele := range inorder {
		if root.Val == ele {
			sliceCutIndex = index
			break
		}
	}
	leftInorder, rightInorder := inorder[0:sliceCutIndex], inorder[sliceCutIndex+1:]

	//根据中序切的集合 分割后续集合 去掉最后一个元素
	leftPostorder, rightPostorder := postorder[0:len(leftInorder)], postorder[len(leftInorder):len(postorder)-1]

	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)
	return root
}

func main() {
	buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
}
