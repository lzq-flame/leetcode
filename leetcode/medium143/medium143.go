/**
2 * @Author: lzq
3 * @Date: 2021/10/4 16:25
4 */

//重排链表
//https://leetcode-cn.com/problems/reorder-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeList(node1, node2 *ListNode) {
	var tmp1, tmp2 *ListNode
	for node1 != nil && node2 != nil {
		tmp1 = node1.Next
		tmp2 = node2.Next

		node1.Next = node2
		node2.Next = tmp1
		node1 = tmp1
		node2 = tmp2
	}
}

func reverseList(head *ListNode) *ListNode {
	var node *ListNode = nil
	curr := head
	for curr != nil {
		pre := curr.Next
		curr.Next = node
		node = curr
		curr = pre
	}
	return node
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	mergeList(l1, l2)
}
