/**
2 * @Author: lzq
3 * @Date: 2021/10/6 10:21
4 */

//删除链表的倒数第N个节点
//https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	preNode := &ListNode{Next: head}
	slow := head
	if fast == nil {
		node := slow.Next
		deleteNode(preNode, slow)
		return node
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
		preNode = preNode.Next
	}
	deleteNode(preNode, slow)
	return head
}

func deleteNode(preNode, node *ListNode) {
	if node == nil {
		return
	}
	preNode.Next = node
	node.Next = nil
}
