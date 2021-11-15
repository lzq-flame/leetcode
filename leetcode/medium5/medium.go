/**
2 * @Author: lzq
3 * @Date: 2021/9/24 15:48
4 */

//最长回文子串
//https://leetcode-cn.com/problems/longest-palindromic-substring/
package main

import "fmt"

func longestPalindrome(s string) string {
	len := len(s)
	if len < 2 {
		return s
	}

	maxLen := 1
	begin := 0

	dp := make([][]bool, len)

	for i := range dp {
		dp[i] = make([]bool, len)
	}
	for i := 0; i < len; i++ {
		dp[i][i] = true
	}
	//总长度
	for L := 2; L <= len; L++ {
		//左边界
		for left := 0; left < len; left++ {
			//右边界
			right := L + left - 1
			if right >= len {
				break
			}
			if s[left] != s[right] {
				dp[left][right] = false
			} else {
				if right-left < 3 {
					dp[left][right] = true
				} else {
					dp[left][right] = dp[left+1][right-1]
				}
			}

			if dp[left][right] && right-left+1 > maxLen {
				maxLen = right - left + 1
				begin = left
			}
		}
	}
	return s[begin : begin+maxLen]
}

func longestPalindrome1(s string) string {
	if len(s) == 1 {
		return s
	}
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right2-left2 > right-left {
			left, right = left2, right2
		}
		if right1-left1 > right-left {
			left, right = left1, right1
		}
	}
	return s[left : right+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {

	}
	return left + 1, right - 1
}

func main() {
	s := "bb"
	palindrome1 := longestPalindrome1(s)
	fmt.Println(palindrome1)
}
