/**
2 * @Author: lzq
3 * @Date: 2021/10/8 17:11
4 */

//排序链表
//https://leetcode-cn.com/problems/sort-list/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	preHead := &ListNode{Next: nil}
	temp := preHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}
	if l1 != nil {
		temp.Next = l1
	}
	if l2 != nil {
		temp.Next = l2
	}
	return preHead.Next
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func main() {
	node3 := &ListNode{3, nil}
	node1 := &ListNode{1, node3}
	node2 := &ListNode{2, node1}
	node4 := &ListNode{4, node2}

	list := sortList(node4)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
