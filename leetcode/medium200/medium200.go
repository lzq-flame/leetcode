/**
2 * @Author: lzq
3 * @Date: 2021/9/25 15:58
4 */
//岛屿数量
//https://leetcode-cn.com/problems/number-of-islands/

package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				ans++
				dfs(grid, i, j)
			}
		}
	}
	return ans
}

func dfs(grid [][]byte, r, c int) {
	if !inArea(grid, r, c) {
		return
	}
	if grid[r][c] != '1' {
		return
	}
	grid[r][c] = '2'
	dfs(grid, r, c+1)
	dfs(grid, r, c-1)
	dfs(grid, r+1, c)
	dfs(grid, r-1, c)

}

func inArea(grid [][]byte, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}
