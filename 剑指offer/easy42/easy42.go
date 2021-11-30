/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 22:37
 */

package main

import "fmt"

//打印出最大和的数组（进阶版）
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 0
	sum := nums[0]
	start, end := 0, 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if sum < count {
			end = i
			sum = count
		}
		if count <= 0 {
			start = i
			count = 0
		}
	}
	fmt.Println(nums[start+1 : end+1])
	return sum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSubArray(nums)
}
