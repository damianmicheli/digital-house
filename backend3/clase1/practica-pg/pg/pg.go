package main

import "fmt"

// import "github.com/gin-gonic/gin"

func main(){

	var nombre string
	nombre = "Damian"

	var horas int = 3
	

	producto, precio := "Jean", 1022.3

	const Minutos = 3
	Minutos1 := 4

	var (
		producto1 = "Curso"
		cantidad = 3
		precio1 = 123
		enStock = true
	)

	var entero int 
	var flotante = float32(entero)

	println(flotante)
	println(entero, "hola mundo")

	fmt.Println("Hola,", nombre)
	fmt.Println(horas)
	fmt.Println(producto, precio, nombre)
	fmt.Println(producto1)
	fmt.Println(Minutos, "minutos")
	fmt.Println(Minutos1, "minutos")
	println(producto1, cantidad, precio1, enStock)
	// fmt.Println(gin.Version)
}