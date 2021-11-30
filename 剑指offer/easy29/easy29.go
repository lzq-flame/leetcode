/**
 * @Author: liuzhiqi
 * @Data: 2021/11/29 10:20
 */

package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, columns-1, 0, rows-1
	index := 0
	ans := make([]int, rows*columns)
	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			ans[index] = matrix[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			ans[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				ans[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- {
				ans[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		bottom--
		top++
	}
	return ans
}
