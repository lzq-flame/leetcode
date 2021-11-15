/**
2 * @Author: lzq
3 * @Date: 2021/10/3 14:11
4 */

//合并k个有序链表
//https://leetcode-cn.com/problems/merge-k-sorted-lists/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}
	mid := (left + right) >> 1
	return mergeTwoLists2(merge(lists, left, mid), merge(lists, mid+1, right))
}

func mergeTwoLists(node1, node2 *ListNode) *ListNode {
	if node1 == nil {
		return node2
	} else if node2 == nil {
		return node1
	} else if node1.Val < node2.Val {
		node1.Next = mergeTwoLists(node1.Next, node2)
		return node1
	} else {
		node2.Next = mergeTwoLists(node2.Next, node1)
		return node2
	}
}

func mergeTwoLists2(node1, node2 *ListNode) *ListNode {
	hair := &ListNode{}
	node := hair
	for node1 != nil && node2 != nil {
		if node1.Val > node2.Val {
			node.Next = node2
			node2 = node2.Next
		} else {
			node.Next = node1
			node1 = node1.Next
		}
		node = node.Next
	}

	if node1 == nil {
		node.Next = node2
	}
	if node2 == nil {
		node.Next = node1
	}
	return hair.Next
}

func main() {
	node1 := &ListNode{Val: 1}
	var node *ListNode = nil

	node3 := mergeTwoLists2(node1, node)
	if node == nil {
		fmt.Println("nil")
	}
	fmt.Println(node3.Val)
}
