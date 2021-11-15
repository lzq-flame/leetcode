/**
2 * @Author: lzq
3 * @Date: 2021/10/20 10:45
4 */

package main

import "fmt"

func findLength(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	ret := 0
	for i := 0; i < n; i++ {
		length := min(m, n-i)
		maxLen := maxLength(nums1, nums2, i, 0, length)
		ret = max(ret, maxLen)
	}

	for i := 0; i < m; i++ {
		length := min(n, m-i)
		maxLen := maxLength(nums1, nums2, 0, i, length)
		ret = max(ret, maxLen)
	}
	return ret
}

func maxLength(A, B []int, addA, addB, length int) int {
	ret, k := 0, 0
	for i := 0; i < length; i++ {
		if A[addA+i] == B[addB+i] {
			k++
		} else {
			k = 0
		}
		ret = max(ret, k)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	a := []int{0, 1, 1, 1, 1}
	b := []int{1, 0, 1, 0, 1}
	length := findLength(a, b)
	fmt.Println(length)
}
