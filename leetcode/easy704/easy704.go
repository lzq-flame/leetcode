/**
2 * @Author: lzq
3 * @Date: 2021/10/3 10:22
4 */

//二分查找
//https://leetcode-cn.com/problems/binary-search/

package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
