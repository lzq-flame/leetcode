/**
 * @Author: liuzhiqi
 * @Data: 2021/11/18 17:14
 */

package main

func sumNums(n int) int {
	ans := 0
	var sum func(i int) bool
	sum = func(i int) bool {
		ans += i
		return i > 0 && sum(i-1)
	}
	sum(n)
	return ans
}
