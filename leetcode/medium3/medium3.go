package main

/**
@time 2021/9/08
*/

//无重复的最长子串
//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	right, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for right+1 < n && m[s[right+1]] == 0 {
			m[s[right+1]]++
			right++
		}
		ans = max(ans, right-i+1)
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
	s := " abcabcbb"
	println(lengthOfLongestSubstring(s))
}
