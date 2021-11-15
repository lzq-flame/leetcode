/**
 * @Author: liuzhiqi
 * @Data: 2021/11/10 16:41
 */

//二叉树最大深度
//https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
