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

func sudokuRowColumnCheck(sudoku [9][9]int) int {
	flag := 0

	//Traverse through each element in 9 X 9 array

	for i := 0; i < 9 && flag != 1; i++ {
		for j := 0; j < 9 && flag != 1; j++ {

			element := sudoku[i][j]

			if element == 0 { //0 means "null" so skipping the element
				continue
			} else {
				if element > 9 {
					flag = 1
					fmt.Println("ERROR - element more than 9\nrow - ", i+1, "\ncolumn - ", j+1)
				}
				//row check
				for r := 0; r < 9 && flag != 1; r++ {
					if r == i { //element should not be compared with itself so skip the part
						continue
					} else {
						if element == sudoku[r][j] {
							flag = 1
							fmt.Println("Element -", element, "\nrow - ", r+1, "\ncolumn - ", j+1)
							break
						}
					}
				}

				//column check
				for c := 0; c < 9 && flag != 1; c++ {
					if c == j { //element should not be compared with itself so skip the part
						continue
					} else {
						if element == sudoku[i][c] {
							flag = 1
							fmt.Println("Element -", element, "\nrow - ", i+1, "\ncolumn - ", c+1)
							break
						}
					}
				}
			}
		}
	}
	return flag
}

func smallSudokuCheck(sudoku [9][9]int, iStart, jStart, iEnd, jEnd int) int {
	var smallSudoku [9][9]int
	for i := iStart; i <= iEnd; i++ {
		for j := jStart; j <= jEnd; j++ {
			smallSudoku[i][j] = sudoku[i][j]
		}
	}

	//To print small sudoku
	/* for i := iStart; i <= iEnd; i++ {
		for j := jStart; j <= jEnd; j++ {
			fmt.Print(smallSudoku[i][j], " ")
		}
		fmt.Println()
	} */

	var flag, count = 0, 0
	for i := iStart; i <= iEnd && flag != 1; i++ {
		for j := jStart; j <= jEnd && flag != 1; j++ {
			element := smallSudoku[i][j]
			count = 0
			if element == 0 { //0 means "null" so skipping the element
				continue
			} else {
				for di := iStart; di <= iEnd && flag != 1; di++ {
					for dj := jStart; dj <= jEnd && flag != 1; dj++ {
						if element == smallSudoku[di][dj] && flag != 1 {
							count++
							if count == 2 {
								flag = 1
								fmt.Println("Number -", element, "\nRow -", di+1, "\nColumn -", dj+1)
								printMatrix(smallSudoku, iStart, jStart, iEnd, jEnd)
								break
							}
						}
					}
				}
			}
		}
	}
	return flag
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
		{0, 0, 0, 0, 8, 0, 0, 2, 9},
	}

	/* printMatrix(sudoku, 0, 0, 8, 8)
	fmt.Println("") */

	v1 := smallSudokuCheck(sudoku, 0, 0, 2, 2)
	v2 := smallSudokuCheck(sudoku, 3, 0, 5, 2)
	v3 := smallSudokuCheck(sudoku, 6, 0, 8, 2)
	v4 := smallSudokuCheck(sudoku, 0, 3, 2, 5)
	v5 := smallSudokuCheck(sudoku, 3, 3, 5, 5)
	v6 := smallSudokuCheck(sudoku, 6, 3, 8, 5)
	v7 := smallSudokuCheck(sudoku, 0, 6, 2, 8)
	v8 := smallSudokuCheck(sudoku, 3, 6, 5, 8)
	v9 := smallSudokuCheck(sudoku, 6, 6, 8, 8)
	sudokuValue := v1 + v2 + v3 + v4 + v5 + v6 + v7 + v8 + v9 //If anyone of them is false then value will be > 0

	if sudokuValue == 0 && sudokuRowColumnCheck(sudoku) == 0 {
		fmt.Println("SUDOKU IS VALID ")
	} else {
		fmt.Println("NOT A VALID SUDOKU")
	}
}
