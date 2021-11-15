/**
2 * @Author: lzq
3 * @Date: 2021/10/4 16:58
4 */

//二叉树右视图

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	stack []interface{}
	size  int
}

func newStack() *Stack {
	return &Stack{stack: make([]interface{}, 0), size: 0}
}

func (s *Stack) push(a interface{}) {
	s.stack = append(s.stack, a)
	s.size++
}

func (s *Stack) pop() interface{} {
	a := s.stack[s.size-1]
	s.stack = s.stack[0 : s.size-1]
	s.size--
	return a
}

func (s *Stack) empty() bool {
	if s.size == 0 {
		return true
	}
	return false
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	rightValue := map[int]int{}
	maxDepth := -1
	depthStack := newStack()
	nodeStack := newStack()
	depthStack.push(0)
	nodeStack.push(root)
	for !nodeStack.empty() {
		node := nodeStack.pop()
		depth := depthStack.pop()
		if node.(*TreeNode) != nil {
			maxDepth = max(depth.(int), maxDepth)
			if _, ok := rightValue[depth.(int)]; !ok {
				rightValue[maxDepth] = node.(*TreeNode).Val
			}
			nodeStack.push(node.(*TreeNode).Left)
			nodeStack.push(node.(*TreeNode).Right)
			depthStack.push(depth.(int) + 1)
			depthStack.push(depth.(int) + 1)
		}
	}
	nums := make([]int, 0)
	for i := 0; i <= maxDepth; i++ {
		nums = append(nums, rightValue[i])
	}
	return nums
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	node5 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 2, Right: node5}
	node3 := &TreeNode{Val: 3, Right: node4}
	node1 := &TreeNode{Val: 1, Left: node2, Right: node3}

	view := rightSideView(node1)
	fmt.Println(view)
}
