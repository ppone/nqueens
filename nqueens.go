//@author: Parag Patel
//this is the first version of nequeens, it has not been optimized yet. Hasn't been fully tested.

package main

import (
	"fmt"
	"os"
	"strconv"
)

var boardsize int = 3

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

	//check left

	for i := x - 1; i >= 0; i-- {
		if board[i][y] == true {
			return false
		}
	}

	//check leftdiagonalup
	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == true {
			return false
		}
	}

	//check leftdiagonaldown
	for i, j := x-1, y+1; i >= 0 && j < boardsize; i, j = i-1, j+1 {
		if board[i][j] == true {
			return false
		}
	}

	return true
}

func newboard(board [][]bool) [][]bool {
	nboard := make([][]bool, boardsize)
	for i, _ := range nboard {
		nboard[i] = make([]bool, boardsize)
	}

	if board != nil {
		copy(nboard, board)
	}

	return nboard
}

func nxnboardandqueensolution(n int) {
	boardsize = n
	fmt.Println(n, "x", n, " board with ", n, " queens, has the following solutions")
	fmt.Println("total of ", nqueens(newboard(nil), n, 0), " solutions")
}

func nqueens(board [][]bool, n int, startx int) int {
	if n == 0 {
		return 1
	}

	counter := 0

	for i := startx; i < boardsize; i++ {
		for j := 0; j < boardsize; j++ {
			if board[i][j] == false && validatePosition(board, i, j) {
				board[i][j] = true
				nboard := newboard(board)
				counter += nqueens(nboard, n-1, i+1)
				board[i][j] = false
			}
		}
	}

	return counter
}

func main() {

	n, _ := strconv.Atoi(os.Args[1])

	for i := 1; i <= n; i++ {
		nxnboardandqueensolution(i)
	}

}
