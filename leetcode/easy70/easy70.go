/**
2 * @Author: lzq
3 * @Date: 2021/10/6 8:26
4 */
//爬楼梯
//https://leetcode-cn.com/problems/climbing-stairs/
package main

func climbStairs(n int) int {
	a, b, c := 1, 2, 0
	if n == 1 {
		return a
	}
	if n == 2 {
		return b
	}
	for i := 2; i < n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}
