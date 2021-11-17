/**
 * @Author: liuzhiqi
 * @Data: 2021/11/15 15:28
 */

package main

func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	return dfs(0, 0, m, n, k, visited)
}

func dfs(i, j, m, n, k int, visited [][]bool) int {
	if i >= m || getSum(i)+getSum(j) > k || j >= n || visited[i][j] {
		return 0
	}
	visited[i][j] = true
	return 1 + dfs(i+1, j, m, n, k, visited) + dfs(i, j+1, m, n, k, visited)
}

func getSum(i int) int {
	sum := i % 10
	tmp := i / 10
	for tmp > 0 {
		sum += tmp % 10
		tmp /= 10
	}
	return sum
}
