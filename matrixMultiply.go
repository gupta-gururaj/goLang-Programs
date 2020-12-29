package main

import "fmt"

func main() {
	//Enter 2 matrixes
	var firstMatrix = [3][2]int{
		{3, 4},
		{4, 5},
		{6, 7},
	}
	//To take values Dynamically
	/* var row1, column1 int
	fmt.Scan(&row1)
	fmt.Scan(&column1)
	fmt.Println("Enter 3 x 2 matrix")
	for i := 0; i < row1; i++ {
		for j := 0; j < column1; j++ {
			fmt.Scan(&firstMatrix[i][j])
		}
	} */

	var secondMatrix = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	//To take values Dynamically
	/* var row1, column1 int
	fmt.Scan(&row2)
	fmt.Scan(&column2)
	fmt.Println("Enter 2 x 3 matrix")
	for i := 0; i < row2; i++ {
		for j := 0; j < column2; j++ {
			fmt.Scan(&secondMatrix[i][j])
		}
	} */

	//Calculating 3rd matrix
	var multipliedMatrix [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			multipliedMatrix[i][j] = 0
			for k := 0; k < 2; k++ {
				multipliedMatrix[i][j] = multipliedMatrix[i][j] + (firstMatrix[i][k] * secondMatrix[k][j])
			}
		}
	}

	fmt.Println("Multiplied matrix")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(multipliedMatrix[i][j], " ")
		}
		fmt.Println()
	}
}
