package main
import "fmt"

func printSmallMatrix(matrix [9][9]int, i1, j1 int) {
	for i := i1; i <= i1+2; i++ {
		for j := j1; j <= j1+2; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func sudokuRowColumnCheck(sudoku [9][9]int) bool {
	//Traverse through each currentValue in 9 X 9 array
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			currentValue := sudoku[i][j]
			if currentValue == 0 { //0 means "null" so skipping the currentValue
				continue
			} else {
				if currentValue > 9 {
					fmt.Println("ERROR - currentValue more than 9\nrow - ", i+1, "\ncolumn - ", j+1)
					return false
				}
				//row check
				for r := 0; r < 9; r++ {
					if r == i { //currentValue should not be compared with itself so skip the part
						continue
					} else {
						if currentValue == sudoku[r][j] {
							fmt.Println("Element -", currentValue, "\nrow - ", r+1, "\ncolumn - ", j+1)
							return false
						}
					}
				}
				//column check
				for c := 0; c < 9; c++ {
					if c == j { //currentValue should not be compared with itself so skip the part
						continue
					} else {
						if currentValue == sudoku[i][c] {
							fmt.Println("Element -", currentValue, "\nrow - ", i+1, "\ncolumn - ", c+1)
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
			currentValue := sudoku[i][j]
			count := 0
			if currentValue == 0 { //0 means "null" so skipping the currentValue
				continue
			} else {
				for idummy := iStart; idummy <= iStart+2; idummy++ {
					for jdummy := jStart; jdummy <= jStart+2; jdummy++ {
						if currentValue == sudoku[idummy][jdummy] {
							count++ // when count = 1, currentValue is matched with itself
							if count == 2 {
								fmt.Println("Number -", currentValue, "\nRow -", idummy+1, "\nColumn -", jdummy+1)
								printSmallMatrix(sudoku, iStart, jStart)
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
