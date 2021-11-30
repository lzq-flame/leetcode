/**
 * @Author: liuzhiqi
 * @Data: 2021/11/19 15:11
 */

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
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i+1])
	root.Right = buildTree(preorder[len(inorder[:i]):], inorder[i+1:])
	return root
}
