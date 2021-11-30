/**
 * @Author: liuzhiqi
 * @Data: 2021/11/23 15:57
 */

package main

func hammingWeight(num uint32) int {
	var i uint32 = 1
	ans := 0
	for num > 0 {
		if i&num == 1 {
			ans++
		}
		num = num >> 1
	}
	return ans
}
