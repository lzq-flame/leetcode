/**
 * @Author: liuzhiqi
 * @Data: 2021/11/11 11:09
 */
//岛屿最大面积
//https://leetcode-cn.com/problems/number-of-islands/solution/dao-yu-lei-wen-ti-de-tong-yong-jie-fa-dfs-bian-li-/
package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				Max := areaMax(grid, i, j)
				fmt.Println(Max)
				ans = max(Max, ans)
			}
		}
	}
	return ans
}

func areaMax(grid [][]int, r, c int) int {
	if !inArea(grid, r, c) {
		return 0
	}
	if grid[r][c] != 1 {
		return 0
	}
	grid[r][c] = 2
	return 1 + areaMax(grid, r, c-1) + areaMax(grid, r, c+1) + areaMax(grid, r-1, c) + areaMax(grid, r+1, c)
}

func inArea(grid [][]int, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{{0, 1, 0, 1, 1}, {1, 1, 1, 0, 0}, {1, 1, 0, 0, 1}, {0, 1, 0, 0, 1}}
	maxAreaOfIsland(grid)
	//fmt.Println(a)
}
