//https://play.golang.org/p/vEDR3zajRpK

package main

import "fmt"

const N int = 9

func printSudoku(arr map[int][]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}

func isSafe(sudoku map[int][]int, row, col, num int) bool {
	//element check in row & col
	for x := 0; x <= 8; x++ {
		if sudoku[row][x] == num {
			return false
		}
	}
	for x := 0; x <= 8; x++ {
		if sudoku[x][col] == num {
			return false
		}
	}

	//check within 3 x 3 array
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if sudoku[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}

func solveSudoku(sudoku map[int][]int, row, col int) bool {
	if row == N-1 && col == N { //it is the last element of sudoku
		return true
	}
	if col == N { //last column of a row so go to next row and column will start from 0
		row++
		col = 0
	}

	if sudoku[row][col] > 0 {
		return solveSudoku(sudoku, row, col+1)
	}

	for num := 0; num <= N; num++ {
		if isSafe(sudoku, row, col, num) {
			sudoku[row][col] = num
			if solveSudoku(sudoku, row, col+1) {
				return true
			}
			sudoku[row][col] = 0
		}
	}
	return false
}

func main() {
	var sudoku = map[int][]int{
		0: {3, 0, 6, 5, 0, 8, 4, 0, 0},
		1: {5, 2, 0, 0, 0, 0, 0, 0, 0},
		2: {0, 8, 7, 0, 0, 0, 0, 3, 1},
		3: {0, 0, 3, 0, 1, 0, 0, 8, 0},
		4: {9, 0, 0, 8, 6, 3, 0, 0, 5},
		5: {0, 5, 0, 0, 9, 0, 6, 0, 0},
		6: {1, 3, 0, 0, 0, 0, 2, 5, 0},
		7: {0, 0, 0, 0, 0, 0, 0, 7, 4},
		8: {0, 0, 5, 2, 0, 6, 3, 0, 0},
	}

	if solveSudoku(sudoku, 0, 0) {
		fmt.Println("Solvable")
		printSudoku(sudoku)
	} else {
		fmt.Println("Not Solvable")
	}
}
