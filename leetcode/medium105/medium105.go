/**
2 * @Author: lzq
3 * @Date: 2021/10/12 19:11
4 */

//重建二叉树
//https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/submissions/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	node := &TreeNode{Val: preorder[0], Left: nil, Right: nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	node.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i+1])
	node.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return node
}
