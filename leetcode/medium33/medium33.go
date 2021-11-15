/**
2 * @Author: lzq
3 * @Date: 2021/9/25 9:44
4 */

//搜索旋转排序数组
//https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

package main

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 0 {
		if target == nums[0] {
			return 0
		} else {
			return -1
		}
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[n-1] >= target && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
