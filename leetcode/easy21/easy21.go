/**
2 * @Author: lzq
3 * @Date: 2021/11/2 15:22
4 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}
}

func mergeTwoLists2(l1, l2 *ListNode) *ListNode {
	hair := &ListNode{}
	pre := hair
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l1 != nil {
		pre.Next = l1
	}
	if l2 != nil {
		pre.Next = l2
	}
	return hair.Next
}

func first(str string) byte {
	ans := map[byte]int{}
	for i := 0; i < len(str); i++ {
		ans[str[i]]++
	}
	for i := 0; i < len(str); i++ {
		if ans[str[i]] == 0 {
			return str[i]
		}
	}
	return ' '
}

func main() {
}
