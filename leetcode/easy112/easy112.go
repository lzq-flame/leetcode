/**
2 * @Author: lzq
3 * @Date: 2021/10/24 15:44
4 */

//路径总和
//https://leetcode-cn.com/problems/path-sum/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var dfs func(node *TreeNode, sum int)
	ans := false
	dfs = func(node *TreeNode, sum int) {
		if node.Left == nil && node.Right == nil {
			sum += node.Val
			if sum == targetSum {
				ans = true
			}
			return
		}
		if node.Left != nil {
			dfs(node.Left, sum+node.Val)
		}
		if node.Right != nil {
			dfs(node.Right, sum+node.Val)
		}
	}
	dfs(root, 0)
	return ans
}
