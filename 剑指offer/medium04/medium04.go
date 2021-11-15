/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 14:40
 */

//二维数组中查找
package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

func main() {
	nums := [][]int{{-1, 3}}
	array := findNumberIn2DArray(nums, 3)
	fmt.Println(array)
}
