/**
2 * @Author: lzq
3 * @Date: 2021/10/16 9:56
4 */

//滑动窗口最大值
//https://leetcode-cn.com/problems/sliding-window-maximum/
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

func maxSlidingWindow1(nums []int, k int) []int {
	ans := make([]int, k)
	q := make([]int, 0)
	var push func(i int)
	push = func(i int) {
		if len(q) > 0 && nums[i] >= nums[q[len(q)]-1] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}
	ans[0] = nums[q[0]]
	for i := k; i < len(nums); i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	window := maxSlidingWindow1(nums, 3)
	fmt.Println(window)
}
