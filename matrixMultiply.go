package main

import "fmt"

func main() {
	//Enter 2 matrixes
	var m1 [3][2]int
	fmt.Println("Enter 3 x 2 matrix")
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&m1[i][j])
		}
	}
	var m2 [2][3]int
	fmt.Println("Enter 2 X 3 matrix")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&m2[i][j])
		}
	}

	//Displaying entered matrix
	/* fmt.Println("Matrix 1  ")
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(m1[i][j])
		}
		fmt.Println()
	}
	fmt.Println("Matrix 2  ")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(m2[i][j])
		}
		fmt.Println()
	} */

	//Calculating 3rd matrix
	var mux [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			mux[i][j] = 0
			for k := 0; k < 2; k++ {
				mux[i][j] = mux[i][j] + (m1[i][k] * m2[k][j])
			}
		}
	}

	fmt.Println("\nMultiplied matrix is : ")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(mux[i][j], " ")
		}
		fmt.Println()
	}
}
