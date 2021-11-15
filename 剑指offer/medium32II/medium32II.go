/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 15:36
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	queue []*TreeNode
	size  int
}

func newQueue() Queue {
	return Queue{queue: make([]*TreeNode, 0), size: 0}
}

func (q *Queue) push(a *TreeNode) {
	q.queue = append(q.queue, a)
	q.size++
}

func (q *Queue) pop() *TreeNode {
	if q.size == 0 {
		return nil
	}
	a := q.queue[0]
	q.queue = q.queue[1:]
	q.size--
	return a
}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := newQueue()
	ans := make([][]int, 0)
	var level func(node *TreeNode)
	level = func(node *TreeNode) {
		q.push(node)
		for !q.isEmpty() {
			currSize := q.size
			num := make([]int, 0)
			for i := 0; i < currSize; i++ {
				a := q.pop()
				num = append(num, a.Val)
				if a.Left != nil {
					q.push(a.Left)
				}
				if a.Right != nil {
					q.push(a.Right)
				}
			}
			ans = append(ans, num)
		}
	}
	level(root)
	return ans
}
