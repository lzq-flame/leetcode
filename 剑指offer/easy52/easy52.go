/**
 * @Author: liuzhiqi
 * @Data: 2021/11/15 10:37
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeA, nodeB := headA, headB
	for nodeA != nil || nodeB != nil {
		if nodeA == nodeB {
			return nodeA
		}
		if nodeA == nil {
			nodeA = headB
		} else {
			nodeA = nodeA.Next
		}
		if nodeB == nil {
			nodeB = headA
		} else {
			nodeB = nodeB.Next
		}
	}
	return nil
}
