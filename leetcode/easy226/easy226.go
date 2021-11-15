/**
2 * @Author: lzq
3 * @Date: 2021/10/27 22:29
4 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	return root
}

func swap(root *TreeNode) *TreeNode {
	if root.Left == nil || root.Right == nil {
		return root
	}
	left := root.Left
	right := root.Right
	root.Left = swap(right)
	root.Right = swap(left)
	return root
}
