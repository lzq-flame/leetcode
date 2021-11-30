/**
 * @Author: liuzhiqi
 * @Data: 2021/11/29 16:03
 */

package main

func singleNumber(nums []int) int {
	one, two := 0, 0
	for _, num := range nums {
		one = one ^ num & ^two
		two = two ^ num & ^one
	}
	return one
}
