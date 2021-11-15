/**
2 * @Author: lzq
3 * @Date: 2021/10/18 11:17
4 */

//二叉树的直径
//https://leetcode-cn.com/problems/diameter-of-binary-tree/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func diameterOfBinaryTree(root *TreeNode) int {
	ans = 1
	maxLength(root)
	return ans - 1
}

func maxLength(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxLength(root.Left)
	right := maxLength(root.Right)
	ans = max(ans, left+right+1)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	node2 := &TreeNode{2, nil, nil}
	node3 := &TreeNode{3, nil, nil}
	node1 := &TreeNode{1, node2, node3}
	tree := diameterOfBinaryTree(node1)
	fmt.Println(tree)
}
