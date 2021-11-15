/**
2 * @Author: lzq
3 * @Date: 2021/10/7 14:52
4 */

//删除排序链表中的重复元素II
//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	preHead := &ListNode{Next: head}
	curr := preHead
	for curr.Next != nil && curr.Next.Next != nil {
		if curr.Next.Val == curr.Next.Next.Val {
			x := curr.Next.Val
			for curr.Next != nil && curr.Next.Next.Val == x {
				curr.Next = curr.Next.Next
			}
		} else {
			curr = curr.Next
		}
	}
	return preHead.Next
}
