/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 19:11
 */

package main

import "fmt"

func fib(n int) int {
	const mod int = 1e9 + 7
	if n < 2 {
		return n
	}
	a0, a1 := 0, 1
	var a2 int
	for i := 2; i <= n; i++ {
		a2 = (a0 + a1) % mod
		a0 = a1
		a1 = a2
	}
	return a2
}

func main() {
	fmt.Println(fib(7))
}
