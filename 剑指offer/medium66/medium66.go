/**
 * @Author: liuzhiqi
 * @Data: 2021/11/23 18:31
 */

package main

func constructArr(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	b := make([]int, len(a))
	b[0] = 1
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1]
	}
	tmp := 1
	for i := len(a) - 2; i >= 0; i-- {
		tmp *= a[i+1]
		b[i] *= tmp
	}
	return b
}
