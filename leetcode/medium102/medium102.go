package main

import (
	"fmt"
)

//二叉树的层次遍历
//https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	value  []interface{}
	length int
}

func NewQueue() *Queue {
	queue := &Queue{
		value:  make([]interface{}, 0),
		length: 0,
	}
	return queue
}

func (q *Queue) Push(a interface{}) {
	q.value = append(q.value, a)
	q.length++
}

func (q *Queue) Pop() interface{} {
	if q.length == 0 {
		return nil
	}
	a := q.value[0]
	q.value = q.value[1:]
	q.length--
	return a

}

func (q *Queue) Empty() bool {
	if q.length == 0 {
		return true
	}
	return false
}
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := NewQueue()
	queue.Push(root)
	for !queue.Empty() {
		currentSize := queue.length
		ans = append(ans, []int{})
		for i := 0; i < currentSize; i++ {
			pop := queue.Pop()
			ans[len(ans)-1] = append(ans[len(ans)-1], pop.(*TreeNode).Val)
			if pop.(*TreeNode).Left != nil {
				queue.Push(pop.(*TreeNode).Left)
			}
			if pop.(*TreeNode).Right != nil {
				queue.Push(pop.(*TreeNode).Right)
			}
		}
	}
	return ans
}
func main() {
	node15 := &TreeNode{
		Left:  nil,
		Right: nil,
		Val:   15,
	}
	node7 := &TreeNode{
		Left:  nil,
		Right: nil,
		Val:   7,
	}
	node20 := &TreeNode{
		Left:  node15,
		Right: node7,
		Val:   20,
	}
	node9 := &TreeNode{
		Left:  nil,
		Right: nil,
		Val:   9,
	}
	node3 := &TreeNode{
		Left:  node9,
		Right: node20,
		Val:   3,
	}

	order := levelOrder(node3)
	for _, m := range order {
		fmt.Println(m)
	}

}
