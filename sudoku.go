package main

import "fmt"

func printMatrix(matrix [9][9]int, i1, j1, i2, j2 int) {
	for i := i1; i <= i2; i++ {
		for j := j1; j <= j2; j++ {
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

func smallSudokuCheck(sudoku [9][9]int, iStart, jStart int) bool {
	var smallSudoku [9][9]int
	for i := iStart; i <= iStart+2; i++ {
		for j := jStart; j <= jStart+2; j++ {
			smallSudoku[i][j] = sudoku[i][j]
		}
	}

	count := 0
	for i := iStart; i <= iStart+2; i++ {
		for j := jStart; j <= jStart+2; j++ {
			element := smallSudoku[i][j]
			count = 0
			if element == 0 { //0 means "null" so skipping the element
				continue
			} else {
				for di := iStart; di <= iStart+2; di++ {
					for dj := jStart; dj <= jStart+2; dj++ {
						if element == smallSudoku[di][dj] {
							count++
							if count == 2 {
								fmt.Println("Number -", element, "\nRow -", di+1, "\nColumn -", dj+1)
								printMatrix(smallSudoku, iStart, jStart, iStart+2, jStart+2)
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

	/* printMatrix(sudoku, 0, 0, 8, 8)
	fmt.Println("") */

	sudokuValue := true
	for i := 0; i <= 6 && sudokuValue; i = i + 3 {
		for j := 0; j <= 6 && sudokuValue; j = j + 3 {
			sudokuValue = smallSudokuCheck(sudoku, i, j)
		}
	}

	/* v1 := smallSudokuCheck(sudoku, 0, 0)
	v2 := smallSudokuCheck(sudoku, 3, 0)
	v3 := smallSudokuCheck(sudoku, 6, 0)
	v4 := smallSudokuCheck(sudoku, 0, 3)
	v5 := smallSudokuCheck(sudoku, 3, 3)
	v6 := smallSudokuCheck(sudoku, 6, 3)
	v7 := smallSudokuCheck(sudoku, 0, 6)
	v8 := smallSudokuCheck(sudoku, 3, 6)
	v9 := smallSudokuCheck(sudoku, 6, 6) */

	if sudokuValue == true && sudokuRowColumnCheck(sudoku) == true {
		fmt.Println("SUDOKU IS VALID ")
	} else {
		fmt.Println("NOT A VALID SUDOKU")
	}
}
