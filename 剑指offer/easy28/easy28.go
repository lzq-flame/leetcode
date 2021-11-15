/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 15:49
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func check(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.Val == node2.Val && check(node1.Left, node2.Right) && check(node1.Right, node2.Left)
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}
