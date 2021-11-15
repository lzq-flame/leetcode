/**
2 * @Author: lzq
3 * @Date: 2021/10/6 17:58
4 */
//两数相加
//https://leetcode-cn.com/problems/add-two-numbers/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || (l1.Next == nil && l1.Val == 0) {
		return l2
	}
	if l2 == nil || (l2.Next == nil && l2.Val == 0) {
		return l1
	}
	preHead := &ListNode{Next: nil}
	head := preHead
	carry, sum, a, b, num := 0, 0, 0, 0, 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		} else {
			a = 0
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		} else {
			b = 0
		}
		sum = a + b + carry
		carry = sum / 10
		num = sum % 10
		node := &ListNode{Next: nil, Val: num}
		head.Next = node
		head = head.Next
	}
	if carry != 0 {
		head.Next = &ListNode{carry, nil}
	}
	return preHead.Next
}

func main() {
	node3 := &ListNode{Next: nil, Val: 3}
	node4 := &ListNode{Next: node3, Val: 4}
	node1 := &ListNode{Next: nil, Val: 1}
	node6 := &ListNode{Next: node1, Val: 6}
	node2 := &ListNode{Next: node4, Val: 2}
	node5 := &ListNode{Next: node6, Val: 5}
	numbers := addTwoNumbers(node2, node5)
	for numbers != nil {
		fmt.Println(numbers.Val)
		numbers = numbers.Next
	}
}
