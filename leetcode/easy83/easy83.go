/**
2 * @Author: lzq
3 * @Date: 2021/10/30 20:59
4 */

//删除链表中的重复元素
//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	curr := head
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}
