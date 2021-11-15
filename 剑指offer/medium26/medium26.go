/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 17:33
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func recur(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}
