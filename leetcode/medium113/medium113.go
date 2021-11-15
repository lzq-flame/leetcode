/**
2 * @Author: lzq
3 * @Date: 2021/10/15 11:21
4 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	path := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, append([]int{}, path...))
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
		path = path[:len(path)-1]
	}
	dfs(root, targetSum)
	return
}

func numsSum(nums []int) int {
	i := 0
	for _, val := range nums {
		i += val
	}
	return i
}

func main() {
	node7 := &TreeNode{7, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node11 := &TreeNode{11, node7, node2}
	node4 := &TreeNode{4, node11, nil}
	node5 := &TreeNode{5, nil, nil}
	node1 := &TreeNode{1, nil, nil}
	node4_2 := &TreeNode{4, node5, node1}
	node13 := &TreeNode{13, nil, nil}
	node8 := &TreeNode{8, node13, node4_2}
	root := &TreeNode{5, node4, node8}

	sum := pathSum(root, 22)
	fmt.Println(sum)
}
