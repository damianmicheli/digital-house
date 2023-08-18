package main

import "fmt"

func main() {
	descuento := 10
	precio := 2000

	resultado := precio - (precio * descuento / 100)

	fmt.Println("Precio:", precio)
	fmt.Println("Descuento:", descuento,"%")
	fmt.Println("Precio final:", resultado)
}