/**
 * @Author: liuzhiqi
 * @Data: 2021/11/23 16:25
 */

package main

func add(a, b int) int {
	for b != 0 {
		c := (a & b) << 1
		a ^= b
		b = c
	}
	return a
}
