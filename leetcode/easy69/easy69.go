/**
2 * @Author: lzq
3 * @Date: 2021/10/6 10:04
4 */
//x的平方根
package main

import "math"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	c, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + c/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}
