/**
2 * @Author: lzq
3 * @Date: 2021/10/12 16:15
4 */

package main

import "fmt"

//最小覆盖子串
//https://leetcode-cn.com/problems/minimum-window-substring/
func minWindow(s string, t string) string {
	flags := make([]bool, 128)
	chars := make([]int, 128)
	for _, value := range t {
		chars[value]++
		flags[value] = true
	}
	cnt, left, min_size, min_left := 0, 0, len(s)+1, 0
	for right := 0; right < len(s); right++ {
		if flags[s[right]] {
			chars[s[right]]--
			if chars[s[right]] >= 0 {
				cnt++
			}
		}
		for cnt == len(t) {
			if right-left+1 < min_size {
				min_left = left
				min_size = right - left + 1
			}
			if flags[s[left]] {
				chars[s[left]]++
				if chars[s[left]] > 0 {
					cnt--
				}
			}
			left++
		}
	}
	if min_size > len(s) {
		return ""
	} else {
		return s[min_left : min_size+min_left]
	}
}
func main() {
	flags := make([]bool, 128)
	chars := make([]int, 128)
	fmt.Println(flags)
	fmt.Println(chars)
}
