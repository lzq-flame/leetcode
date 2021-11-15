/**
2 * @Author: lzq
3 * @Date: 2021/9/21 10:38
4 */

//二叉树的锯齿形层序遍历
//https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	value  []*TreeNode
	length int
}

func NewQueue() *Queue {
	queue := &Queue{
		value: make([]*TreeNode, 0),
	}
	return queue
}

func (q *Queue) Push(node *TreeNode) {
	q.value = append(q.value, node)
	q.length++
}

func (q *Queue) Pop() *TreeNode {
	if q.length == 0 {
		return nil
	}
	node := q.value[0]
	q.value = q.value[1:]
	q.length--
	return node
}

func (q *Queue) Empty() bool {
	if q.length == 0 {
		return true
	}
	return false
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	queue := NewQueue()
	queue.Push(root)
	target := 1
	for !queue.Empty() {
		currentSize := queue.length
		ans = append(ans, []int{})
		for i := 0; i < currentSize; i++ {
			node := queue.Pop()
			ans[len(ans)-1] = append(ans[len(ans)-1], node.Val)
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
		}
		if target%2 == 0 {
			nodeVal := ans[len(ans)-1]
			for i, n := 0, len(nodeVal); i < n/2; i++ {
				nodeVal[i], nodeVal[n-i-1] = nodeVal[n-i-1], nodeVal[i]
			}
		}
		target++
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

	order := zigzagLevelOrder(node3)
	for _, m := range order {
		fmt.Println(m)
	}

}
