/**
 * @Author: liuzhiqi
 * @Data: 2021/11/18 11:39
 */

//单词搜索
package main

func exist(board [][]byte, word string) bool {
	isTraverse := make([][]bool, len(board))
	for i := 0; i < len(isTraverse); i++ {
		isTraverse[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, 0, isTraverse) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, index int, isBoard [][]bool) bool {
	if index == len(word) {
		return true
	}
	if !isArea(board, i, j) || isBoard[i][j] {
		return false
	}
	if board[i][j] == word[index] && !isBoard[i][j] {
		isBoard[i][j] = true
		res := dfs(board, word, i+1, j, index+1, isBoard) ||
			dfs(board, word, i-1, j, index+1, isBoard) ||
			dfs(board, word, i, j+1, index+1, isBoard) ||
			dfs(board, word, i, j-1, index+1, isBoard)
		isBoard[i][j] = false
		return res
	}
	return false
}
func isArea(board [][]byte, i, j int) bool {
	return i >= 0 && i < len(board) && j >= 0 && j < len(board[0])
}
