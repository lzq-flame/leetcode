/**
 * @Author: liuzhiqi
 * @Data: 2021/11/12 10:50
 */
//岛屿周长
package main

func islandPerimeter(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return dfs(grid, i, j)
			}
		}
	}
	return 0
}

func dfs(grid [][]int, r, c int) int {
	if !isArea(grid, r, c) {
		return 1
	}
	if grid[r][c] == 0 {
		return 1
	}
	if grid[r][c] != 1 {
		return 0
	}
	grid[r][c] = 2
	return dfs(grid, r-1, c) + dfs(grid, r+1, c) + dfs(grid, r, c-1) + dfs(grid, r, c+1)

}

func isArea(grid [][]int, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}
