/**
2 * @Author: lzq
3 * @Date: 2021/10/7 10:55
4 */

//区间合并
//https://leetcode-cn.com/problems/merge-intervals/
package main

import (
	"fmt"
	"sort"
)

type interval [][]int

func (i interval) Len() int {
	return len(i)
}

func (i interval) Less(a, b int) bool {
	return i[a][0] < i[b][0]
}

func (i interval) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func merge(intervals [][]int) [][]int {
	nums := make(interval, 0)
	nums = intervals
	sort.Sort(nums)
	ans := make([][]int, 0)
	ans = append(ans, nums[0])
	for i := 1; i < len(nums); i++ {
		two := mergeTwo(ans[len(ans)-1], nums[i])
		if two == nil {
			ans = append(ans, nums[i])
		} else {
			ans[len(ans)-1] = two
		}
	}
	return ans
}

func mergeTwo(a, b []int) []int {
	if b[0] > a[1] {
		return nil
	} else {
		if b[1] < a[1] {
			return a
		} else {
			return []int{a[0], b[1]}
		}
	}
}
func main() {
	nums := make([][]int, 0)
	nums = append(nums, []int{1, 3})
	nums = append(nums, []int{2, 6})
	nums = append(nums, []int{8, 10})
	nums = append(nums, []int{15, 18})
	i := merge(nums)
	fmt.Println(i)
}
