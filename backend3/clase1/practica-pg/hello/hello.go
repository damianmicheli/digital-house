package main

import (
	"fmt"
	"ejemplo/saludo"
)

func main() {
	mensaje := saludo.Hola("Damian")
	fmt.Println(mensaje)
}