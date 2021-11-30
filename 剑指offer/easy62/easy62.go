/**
 * @Author: liuzhiqi
 * @Data: 2021/11/28 18:05
 */

//圆圈中最后剩下的数字

package main

import "fmt"

func lastRemaining(n int, m int) int {
	ans := 0
	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}

func main() {
	fmt.Println(lastRemaining(10, 17))
}
