/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 13:36
 */

//在排序数组中查找数字
package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right + 1) >> 1
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	//for left < right {
	//	mid := (left + right) / 2
	//	if nums[mid] <= target {
	//		right = mid
	//	} else {
	//		left = mid + 1
	//	}
	//}
	ans := 0
	for left >= 0 && nums[left] == target {
		ans++
		left--
	}
	return ans

}

func main() {
	nums := []int{1, 2, 3}
	i := search(nums, 1)
	fmt.Println(i)
}
