/**
2 * @Author: lzq
3 * @Date: 2021/9/9 14:36
4 */

//链表反转
//https://leetcode-cn.com/problems/reverse-linked-list/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	node1 := &ListNode{
		1,
		nil,
	}
	node2 := &ListNode{
		2,
		nil,
	}
	node3 := &ListNode{
		3,
		nil,
	}
	node4 := &ListNode{
		4,
		nil,
	}
	node5 := &ListNode{
		5,
		nil,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	for head := node1; head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
	fmt.Println("after reverse")
	reverseList(node1)
	for head := node5; head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
}
