/**
 * @Author: liuzhiqi
 * @Data: 2021/11/15 09:30
 */

package main

func lengthOfLongestSubstring(s string) int {
	hash := map[byte]int{}
	length := len(s)
	right, ans := -1, 0
	for left := 0; left < length; left++ {
		if left != 0 {
			delete(hash, s[left-1])
		}
		for right+1 < length && hash[s[right+1]] == 0 {
			hash[s[right+1]]++
			right++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
