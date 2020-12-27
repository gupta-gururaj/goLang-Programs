package main

import "fmt"

func main() {
	//Enter matrix size
	var r, c int
	fmt.Println("Enter no. of rows and columns of your matrix")
	fmt.Scan(&r)
	fmt.Scan(&c)
	var m1 [10][10]int
	fmt.Println("Enter your ", r, " X ", c, "matrix")
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Scan(&m1[i][j])
		}
	}

	//Transpose conversion
	var temp int
	for j := 0; j < c; j++ {
		for i := j; i < r; i++ {
			if i != j {
				temp = m1[i][j]
				m1[i][j] = m1[j][i]
				m1[j][i] = temp
			} else {
				continue
			}
		}
	}

	fmt.Println("Transpose Matrix")
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Print(m1[i][j], " ")
		}
		fmt.Println()
	}
}
