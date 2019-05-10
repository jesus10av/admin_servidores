package main

import (
	"fmt"
	"math/rand"
)

func generateMatrix(size int) [][]int {
	matrix := make([][]int, size) //slice for dinamic size array, allocates 0s
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	} //end for

	//rand.Intn(100)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] = rand.Intn(100)
		} //end for
	} //end for

	fmt.Printf("%v", matrix)
	return matrix
} //end generateVector

func generateVector(size int) []int {
	generatedVector := make([]int, size) //slice for dinamic size array, allocates 0s
	for i := 0; i < size; i++ {
		generatedVector[i] = rand.Intn(100)
	} //end for
	fmt.Printf("%v", generatedVector)
	return generatedVector
} //end generateVector

func multiplication(matrix [][]int, vector []int) []int {
	result := make([]int, len(vector))
	for i := 0; i < len(vector); i++ {
		calculatedValue := 0
		for j := 0; j < len(vector); j++ {
			calculatedValue = calculatedValue + (matrix[i][j] * vector[j])
		} //end for
		result[i] = calculatedValue
	} //end for
	fmt.Printf("%v", result)
	return result
} //end multiplication

func main() {
	aMatrix := generateMatrix(5)
	aVector := generateVector(5)
	fmt.Printf("\nRESULT:")
	multiplication(aMatrix, aVector)
} //end main