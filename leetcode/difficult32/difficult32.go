/**
2 * @Author: lzq
3 * @Date: 2021/10/26 11:07
4 */

package main

import "fmt"

//最长有效括号
//https://leetcode-cn.com/problems/longest-valid-parentheses/
func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	stack = append(stack, -1)
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				ans = max(ans, i-stack[len(stack)-1])
			}
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	parentheses := longestValidParentheses(")()())")
	fmt.Println(parentheses)
}
