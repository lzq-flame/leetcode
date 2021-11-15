/**
 * @Author: liuzhiqi
 * @Data: 2021/11/15 10:53
 */

//调整数组使奇数位于偶数前面
package main

func exchange(nums []int) []int {
	for i, n := 0, len(nums)-1; i < len(nums); i++ {
		for nums[i]%2 == 0 && i < n {
			if nums[n]%2 == 1 {
				nums[i], nums[n] = nums[n], nums[i]
			}
			n--
		}
	}
	return nums
}
