/**
2 * @Author: lzq
3 * @Date: 2021/10/6 8:48
4 */

//链表中倒数第k个节点
//https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func main() {

	node2 := &ListNode{Next: nil, Val: 2}
	node1 := &ListNode{Next: node2, Val: 1}
	end := getKthFromEnd(node1, 1)
	fmt.Println(end.Val)

}
