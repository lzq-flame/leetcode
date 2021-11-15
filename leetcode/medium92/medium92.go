/**
2 * @Author: lzq
3 * @Date: 2021/9/29 9:50
4 */
//反转链表II

//https://leetcode-cn.com/problems/reverse-linked-list-ii/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	rightNode := pre

	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}
	leftNode := pre.Next
	curr := rightNode.Next
	for curr != rightNode {
		nex := leftNode.Next
		leftNode.Next = curr
		curr = leftNode
		leftNode = nex
	}
	pre.Next = rightNode
	return dummyNode.Next
}
