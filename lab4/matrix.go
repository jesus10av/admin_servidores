    
package main

import (
	"fmt"
)

func generateMatrix(size int) [][]int {
	matrix := make([][]int, size) //slice for dinamic size array, allocates 0s
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	} //end for
	fmt.Printf("%v", matrix)
	return matrix
} //end generateVector

func main() {
	generateMatrix(5)
} //end main