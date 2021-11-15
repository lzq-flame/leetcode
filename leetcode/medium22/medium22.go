/**
2 * @Author: lzq
3 * @Date: 2021/10/27 16:17
4 */

package main

import "fmt"

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	ans := make([]string, 0)
	res := map[string]int{}
	var tmp string
	for _, value := range generateParenthesis(n - 1) {
		for i := 0; i < 2*(n-1); i++ {
			tmp = value[0:i] + "()" + value[i:]
			if res[tmp] == 0 {
				res[tmp] = 1
				ans = append(ans, tmp)
			}
		}
	}
	return ans
}

func main() {
	parenthesis := generateParenthesis(3)
	fmt.Println(parenthesis)
}
