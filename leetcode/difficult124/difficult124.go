/**
2 * @Author: lzq
3 * @Date: 2021/10/6 9:30
4 */

//二叉树中最大路径和
//https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	var GainSum func(node *TreeNode) int
	GainSum = func(node *TreeNode) int {
		if root == nil {
			return 0
		}
		LeftGain := max(maxPathSum(node.Left), 0)
		RightGain := max(maxPathSum(node.Right), 0)
		path := LeftGain + RightGain + root.Val
		maxSum = max(path, maxSum)
		return max(LeftGain, RightGain) + node.Val
	}
	GainSum(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
