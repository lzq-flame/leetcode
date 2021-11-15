/**
2 * @Author: lzq
3 * @Date: 2021/10/25 16:03
4 */

package main

func majorityElement(nums []int) int {
	count := 0
	var candidate int
	for _, value := range nums {
		if count == 0 {
			candidate = value
		}
		if value == candidate {
			count += 1
		} else {
			count += -1
		}
	}
	return candidate
}
