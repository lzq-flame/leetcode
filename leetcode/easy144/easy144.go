/**
2 * @Author: lzq
3 * @Date: 2021/10/7 10:42
4 */

//二叉树的前序遍历
//https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {

	ans := make([]int, 0)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return ans
}
