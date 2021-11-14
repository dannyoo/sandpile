package main
import (
	"fmt"
	"strconv"
)
func printMatrix(s2 [][]int) {
	n := len(s2)
	// columns in
	m := len(s2[0])
	for i := 0; i < n; i++ { //rows
		fmt.Print("[ ")
		for j := 0; j < m; j++ { //colums
			fmt.Print(strconv.Itoa(s2[i][j]) + " ")
		}
		fmt.Print("]\n")
	}
}