/**
2 * @Author: lzq
3 * @Date: 2021/9/23 18:01
4 */

package main

import "fmt"

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return true

}

func main() {
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	if pairs['}'] > 0 {
		fmt.Println("123312")
	}
}
