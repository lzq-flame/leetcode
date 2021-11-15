/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 09:36
 */

//剑指 Offer 06. 从尾到头打印链表
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	ans := make([]int, 0)
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	reverse(ans)
	return ans
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < len(nums)/2; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
