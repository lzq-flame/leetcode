/**
2 * @Author: lzq
3 * @Date: 2021/9/21 18:01
4 */

package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p3 := m-1, n-1, m+n-1
	for ; p1 >= 0 || p2 >= 0; p3-- {
		var curr int
		if p1 == -1 {
			curr = nums2[p2]
			p2--
		} else if p2 == -1 {
			curr = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			curr = nums1[p1]
			p1--
		} else {
			curr = nums2[p2]
			p2--
		}
		nums1[p3] = curr
	}
}

func exchange(a, b []int) {
	a = b
	fmt.Println(a)
}

func main() {
	nums1 := []int{0}
	nums2 := []int{1}
	exchange(nums1, nums2)
	fmt.Println(nums1)
}
