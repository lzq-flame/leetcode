/**
 * @Author: liuzhiqi
 * @Data: 2021/11/15 18:24
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var prenode func(node *TreeNode)
	ans := 0
	prenode = func(node *TreeNode) {
		if node == nil {
			return
		}
		prenode(node.Right)
		if k == 0 {
			return
		}
		k--
		if k == 0 {
			ans = node.Val
		}
		prenode(node.Left)
	}
	prenode(root)
	return ans
}
