/**
2 * @Author: lzq
3 * @Date: 2021/10/16 17:41
4 */

//平衡二叉树
//https://leetcode-cn.com/problems/balanced-binary-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0

}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
