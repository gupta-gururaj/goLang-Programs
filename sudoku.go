package main

import "fmt"

func printMatrix(matrix [9][9]int, i1, j1 int) {
	for i := i1; i <= i1+2; i++ {
		for j := j1; j <= j1+2; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func sudokuRowColumnCheck(sudoku [9][9]int) bool {
	//Traverse through each element in 9 X 9 array

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			element := sudoku[i][j]

			if element == 0 { //0 means "null" so skipping the element
				continue
			} else {
				if element > 9 {
					fmt.Println("ERROR - element more than 9\nrow - ", i+1, "\ncolumn - ", j+1)
					return false
				}
				//row check
				for r := 0; r < 9; r++ {
					if r == i { //element should not be compared with itself so skip the part
						continue
					} else {
						if element == sudoku[r][j] {
							fmt.Println("Element -", element, "\nrow - ", r+1, "\ncolumn - ", j+1)
							return false
						}
					}
				}

				//column check
				for c := 0; c < 9; c++ {
					if c == j { //element should not be compared with itself so skip the part
						continue
					} else {
						if element == sudoku[i][c] {
							fmt.Println("Element -", element, "\nrow - ", i+1, "\ncolumn - ", c+1)
							return false
						}
					}
				}
			}
		}
	}
	return true
}

func sudokuCheck(sudoku [9][9]int, iStart, jStart int) bool {
	for i := iStart; i <= iStart+2; i++ {
		for j := jStart; j <= jStart+2; j++ {
			element := sudoku[i][j]
			count := 0
			if element == 0 { //0 means "null" so skipping the element
				continue
			} else {
				for di := iStart; di <= iStart+2; di++ {
					for dj := jStart; dj <= jStart+2; dj++ {
						if element == sudoku[di][dj] {
							count++ // when count = 1, element is matched with itself
							if count == 2 {
								fmt.Println("Number -", element, "\nRow -", di+1, "\nColumn -", dj+1)
								printMatrix(sudoku, iStart, jStart)
								return false
							}
						}
					}
				}
			}
		}
	}
	return true
}

func main() {
	var sudoku = [9][9]int{
		{5, 4, 0, 0, 7, 0, 0, 0, 0}, //0 means null
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 7, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	allSmallSudokus := true
	for i := 0; i <= 6 && allSmallSudokus; i = i + 3 {
		for j := 0; j <= 6 && allSmallSudokus; j = j + 3 {
			allSmallSudokus = sudokuCheck(sudoku, j, i)
		}
	}

	/* v1 := sudokuCheck(sudoku, 0, 0)
	v2 := sudokuCheck(sudoku, 3, 0)
	v3 := sudokuCheck(sudoku, 6, 0)
	v4 := sudokuCheck(sudoku, 0, 3)
	v5 := sudokuCheck(sudoku, 3, 3)
	v6 := sudokuCheck(sudoku, 6, 3)
	v7 := sudokuCheck(sudoku, 0, 6)
	v8 := sudokuCheck(sudoku, 3, 6)
	v9 := sudokuCheck(sudoku, 6, 6) */

	if allSmallSudokus == true && sudokuRowColumnCheck(sudoku) == true {
		fmt.Println("SUDOKU IS VALID ")
	} else {
		fmt.Println("NOT A VALID SUDOKU")
	}
}
