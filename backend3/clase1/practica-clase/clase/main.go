package main

import "fmt"

func main() {

	word := "random"

	//imprimir la longintud de la palabra
	fmt.Println("La longitud de word es", len(word))

	//imprimir cada una de las letras
	for _, value := range word {
		fmt.Println(string(value))
	}
}