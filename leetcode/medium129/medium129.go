/**
2 * @Author: lzq
3 * @Date: 2021/10/13 17:08
4 */
//求根节点到叶节点数字之和
//https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var sum func(node *TreeNode, num int)
	ans := 0
	sum = func(node *TreeNode, num int) {
		if node.Left == nil && node.Right == nil {
			num = num*10 + node.Val
			ans += num
			return
		}
		if node.Left != nil {
			sum(node.Left, num*10+node.Val)
		}
		if node.Right != nil {
			sum(node.Right, num*10+node.Val)
		}
	}
	sum(root, 0)
	return ans
}

func main() {

	node5 := &TreeNode{5, nil, nil}
	node1 := &TreeNode{1, nil, nil}
	node9 := &TreeNode{9, node5, node1}
	node0 := &TreeNode{0, nil, nil}
	node4 := &TreeNode{4, node9, node0}
	numbers := sumNumbers(node4)
	fmt.Println(numbers)
}
