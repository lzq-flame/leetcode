/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 18:46
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	return swap(root)
}

func swap(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := root.Left
	right := root.Right
	root.Left = swap(right)
	root.Right = swap(left)
	return root
}
