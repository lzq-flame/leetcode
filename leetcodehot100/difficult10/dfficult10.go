/**
 * @Author: liuzhiqi
 * @Data: 2021/11/17 11:18
 */

package main

func isMatch(s string, p string) bool {
	if len(s) == 0 || len(p) == 0 {
		return false
	}
	return matchCore(s, p, 0, 0)
}
func matchCore(s, p string, indexS, indexP int) bool {
	if indexS >= len(s) && indexP >= len(p) {
		return true
	}
	if indexS < len(s) && indexP >= len(p) {
		return false
	}
	if indexP+1 < len(p) && p[indexP+1] == '*' {
		if indexS < len(s) && ((p[indexP] == s[indexS]) || (p[indexP] == '.' && indexS < len(s))) {
			return matchCore(s, p, indexS+1, indexP+2) ||
				matchCore(s, p, indexS+1, indexP) ||
				matchCore(s, p, indexS, indexP+2)
		} else {
			return matchCore(s, p, indexS, indexP+2)
		}
	}
	if indexP < len(p) && indexS < len(s) && ((p[indexP] == s[indexS]) || (p[indexP] == '.' && indexS < len(s))) {
		return matchCore(s, p, indexS+1, indexP+1)
	}
	return false
}

//动态规划版
func isMatch1(s string, p string) bool {
	m, n := len(s), len(p)
	if m == 0 || n == 0 {
		return false
	}
	var match func(i, j int) bool
	match = func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}
	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				if match(i, j-1) {
					f[i][j] = f[i-1][j] || f[i][j-2]
				} else {
					f[i][j] = f[i][j-2]
				}
			} else {
				if match(i, j) {
					f[i][j] = f[i-1][j-1]
				}
			}
		}
	}
	return f[m][n]
}
