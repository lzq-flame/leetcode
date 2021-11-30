/**
 * @Author: liuzhiqi
 * @Data: 2021/11/23 16:49
 */

package main

func singleNumbers(nums []int) []int {
	ret := 0
	for _, value := range nums {
		ret ^= value
	}
	div := 1
	for div&ret == 0 {
		div = div << 1
	}
	a, b := 0, 0
	for _, value := range nums {
		if div&value == 0 {
			a ^= value
		} else {
			b ^= value
		}
	}
	return []int{a, b}
}
