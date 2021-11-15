/**
2 * @Author: lzq
3 * @Date: 2021/9/17 16:00
4 */

//相交链表
//https://leetcode-cn.com/problems/intersection-of-two-linked-lists/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	A := headA
	B := headB
	for A != B {
		if A == nil {
			A = headB
		} else {
			A = A.Next
		}
		if B == nil {
			B = headA
		} else {
			B = B.Next
		}
	}
	return A
}

func main() {

}
