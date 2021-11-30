/**
 * @Author: liuzhiqi
 * @Data: 2021/11/23 18:06
 */

package main

func majorityElement(nums []int) int {
	var candidate int
	count := 0
	for _, value := range nums {
		if count == 0 {
			candidate = value
		}
		if candidate == value {
			count++
		} else {
			count--
		}
	}
	return candidate
}
