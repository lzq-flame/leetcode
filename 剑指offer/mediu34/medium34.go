/**
 * @Author: liuzhiqi
 * @Data: 2021/11/15 17:20
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(node *TreeNode, num int)
	dfs = func(node *TreeNode, num int) {
		if node == nil {
			return
		}
		num -= node.Val
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && num == 0 {
			ans = append(ans, append([]int(nil), path...))
		}
		dfs(node.Left, num)
		dfs(node.Right, num)
		path = path[:len(path)-1]
	}
	dfs(root, target)
	return ans
}
