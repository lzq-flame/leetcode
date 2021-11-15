/**
2 * @Author: lzq
3 * @Date: 2021/10/30 11:13
4 */

//34. 在排序数组中查找元素的第一个和最后一个位置
//https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
package main

import "fmt"

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if nums[left] != target {
		return ans
	}
	ans[0] = left
	left, right = 0, len(nums)-1
	for left < right {
		mid := (left + right + 1) >> 1
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	if nums[left] != target {
		return ans
	}
	ans[1] = left
	return ans
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	ints := searchRange(nums, 8)
	fmt.Println(ints)
}
