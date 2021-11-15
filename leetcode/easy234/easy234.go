/**
2 * @Author: lzq
3 * @Date: 2021/10/24 15:08
4 */

//回文链表
//https://leetcode-cn.com/problems/palindrome-linked-list/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	arr := make([]int, 0)
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	left, right := head, slow
	for right != nil {
		arr = append(arr, right.Val)
		right = right.Next
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != left.Val {
			return false
		}
		left = left.Next
	}
	return true
}
func main() {
	node11 := &ListNode{1, nil}
	node12 := &ListNode{2, node11}
	node2 := &ListNode{2, node12}
	node1 := &ListNode{1, node2}
	palindrome := isPalindrome(node1)
	fmt.Println(palindrome)
}
