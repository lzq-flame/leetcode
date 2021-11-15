/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 22:37
 */

package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 0
	sum := nums[0]
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if sum < count {
			sum = count
		}
		if count <= 0 {
			count = 0
		}
	}
	return sum
}
