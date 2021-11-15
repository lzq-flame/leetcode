/**
 * @Author: liuzhiqi
 * @Data: 2021/11/11 15:06
 */

//填海造陆问题
//
package main

import (
	"fmt"
)

func largestIsland(grid [][]int) int {
	var maxarea []int
	if len(grid)*len(grid[0]) <= 3 {
		maxarea = make([]int, 3)
	} else {
		maxarea = make([]int, len(grid)*len(grid[0]))
	}
	tag := 2
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				a := dfs(grid, tag, i, j)
				maxarea[tag] = a
				tag++
			}
		}
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			temp := map[int]bool{}
			if grid[i][j] == 0 {
				up, down, left, right := 0, 0, 0, 0
				if isArea(grid, i-1, j) {
					up = maxarea[grid[i-1][j]]
					temp[grid[i-1][j]] = true
				}
				if isArea(grid, i+1, j) && !temp[grid[i+1][j]] {
					down = maxarea[grid[i+1][j]]
					temp[grid[i+1][j]] = true
				}
				if isArea(grid, i, j-1) && !temp[grid[i][j-1]] {
					left = maxarea[grid[i][j-1]]
					temp[grid[i][j-1]] = true
				}
				if isArea(grid, i, j+1) && !temp[grid[i][j+1]] {
					right = maxarea[grid[i][j+1]]
					temp[grid[i][j+1]] = true
				}
				Max := 1 + up + down + left + right
				ans = max(Max, ans)
			}
		}
	}
	if ans == 0 {
		return len(grid) * len(grid[0])
	}
	return ans
}

func dfs(grid [][]int, tag, r, c int) int {
	if !isArea(grid, r, c) {
		return 0
	}
	if grid[r][c] != 1 {
		return 0
	}
	grid[r][c] = tag
	return 1 + dfs(grid, tag, r, c-1) + dfs(grid, tag, r, c+1) + dfs(grid, tag, r-1, c) + dfs(grid, tag, r+1, c)
}

func isArea(grid [][]int, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 0, 0},
		{0, 1, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 0, 0},
		{0, 1, 1, 1, 1, 0, 0},
	}
	a := largestIsland(grid)
	fmt.Println(a)
}
