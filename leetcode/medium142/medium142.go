/**
2 * @Author: lzq
3 * @Date: 2021/9/25 16:46
4 */

//环形链表II
//https://leetcode-cn.com/problems/linked-list-cycle-ii/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != fast {
				p = p.Next
				fast = fast.Next
			}
			return p
		}
	}
	return nil
}
