/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 14:18
 */

//0～n-1 缺失的数字

package main

func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
