//@author: Parag Patel
//this is the first version of nQueens.  It runs fast, but it still not the fastest solution in Go.

package main

import (
	"fmt"
	"os"
	"strconv"
)

var boardSize int = 3

func printBoard(board [][]bool) {
	fmt.Println("<==SOLUTION==>")
	for i, _ := range board {
		for j, _ := range board[i] {
			if board[i][j] == false {
				fmt.Printf("B ")
			} else {
				fmt.Printf("Q ")
			}
		}
		fmt.Printf("\n")
	}
}

func validatePosition(board [][]bool, x int, y int) bool {

	//check up
	for i := x - 1; i >= 0; i-- {
		if board[i][y] == true {
			return false
		}
	}

	//check leftDiagonalUp
	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == true {
			return false
		}
	}

	//check rightDiagonalUp
	for i, j := x-1, y+1; i >= 0 && j < boardSize; i, j = i-1, j+1 {
		if board[i][j] == true {
			return false
		}
	}

	return true
}

func newBoard(board [][]bool) [][]bool {
	nBoard := make([][]bool, boardSize)
	for i, _ := range nBoard {
		nBoard[i] = make([]bool, boardSize)
	}

	if board != nil {
		copy(nBoard, board)
	}

	return nBoard
}

func nxnBoardAndQueenSolution(n int) {
	boardSize = n
	fmt.Println(n, "x", n, " board with ", n, " queens, has the following solutions")
	fmt.Println("total of ", nQueens(newBoard(nil), 0), " solutions")
}

func nQueens(board [][]bool, n int) int {
	if n == boardSize {
		return 1
	}

	counter := 0

	for j := 0; j < boardSize; j++ {
		if board[n][j] == false && validatePosition(board, n, j) {
			board[n][j] = true
			nBoard := newBoard(board)
			counter += nQueens(nBoard, n+1)
			board[n][j] = false
		}

	}

	return counter
}

func main() {

	n, _ := strconv.Atoi(os.Args[1])

	//for i := 1; i <= n; i++ {
	nxnBoardAndQueenSolution(n)
	//}

}
