package leetcode79

import (
	"strconv"
	"strings"
)

var directions = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

var m, n int

func exist(board [][]byte, word string) bool {
	if word == "" || len(word) == 0 {
		return true
	}
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	haveVisited := make([][]bool, m*n)
	for i := 0; i < m; i++ {
		haveVisited[i] = make([]bool, n)
	}
	m, n = len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if traverse(0, i, j, &haveVisited, board, word) {
				return true
			}
		}
	}
	return false
}

func traverse(currentLen, r, c int, visited *[][]bool, board [][]byte, word string) bool {
	if currentLen == len(word) {
		return true
	}

	if r >= m || c >= n || r < 0 || c < 0 || board[r][c] != word[currentLen] || (*visited)[r][c] {
		return false
	}

	(*visited)[r][c] = true
	for _, d := range directions {
		if traverse(currentLen, r+d[0], c+d[1], visited, board, word) {
			return false
		}
	}
	(*visited)[r][c] = false
	return false
}

func main(){
	strings.Join()

	strconv.Itoa()
}