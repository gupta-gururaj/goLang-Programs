package main

import (
	"fmt"
)

func main() {
	//Enter orgMatrix
	var row, column int
	var orgMatrix [10][10]int

	fmt.Println("Enter no. of rows and columns of your orgMatrix")
	fmt.Scan(&row)
	fmt.Scan(&column)
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			fmt.Scan(&orgMatrix[i][j])
		}
	}

	//taking transpose of the Matrix
	var transposedMatrix [10][10]int
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			transposedMatrix[j][i] = orgMatrix[i][j]
		}
	}

	fmt.Println("Transpose Matrix")
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			fmt.Print(transposedMatrix[i][j], " ")
		}
		fmt.Println()
	}
}
