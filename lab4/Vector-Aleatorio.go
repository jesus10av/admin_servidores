/*
*Programa que crea un vector de longitud N de integer con 0s
*/
package main

import "fmt"
import "math/rand"

//FUNCIONES
func crearVectorAleatorio(size int, maxRandom int) []int {
	array:= make([]int, size)

	for i := 0; i < size; i++{
		array[i] = rand.Intn(maxRandom)
	}//End for 

	return array
}

//Inicio del programa
func main() {
	resultado := crearVectorAleatorio(200,20)
	//Print vector
	fmt.Printf("%v",resultado)
}