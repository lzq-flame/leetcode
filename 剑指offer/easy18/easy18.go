/**
 * @Author: liuzhiqi
 * @Data: 2021/11/15 09:47
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	node := &ListNode{Next: head}
	pre := node
	next := pre.Next
	for next != nil {
		if next.Val == val {
			pre.Next = next.Next
		}
		pre = next
		next = next.Next
	}
	return node.Next
}
