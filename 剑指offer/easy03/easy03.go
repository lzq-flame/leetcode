/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 11:18
 */

//数组中重复的数字
package main

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[nums[i]] != nums[i] {
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			} else {
				return nums[i]
			}
		}
	}
	return 0
}
