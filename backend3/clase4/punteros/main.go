package main

import "fmt"

func Incrementar(v *int) {
	*v++
}

func main() {
/*
	var p *int //crear puntero del tipo int

	var p1 *float64 // crear puntero del tipo float

	var p2 = new(float64) // crear otro puntero float con new

	var v float64 // crear variable float

	p3 := &v // crear puntero float usando operador de direccion
*/
	var v1 int = 50

	fmt.Println("La direccion de memoria de la variable v1 es:", &v1)
	fmt.Println("El valor contenido en la variable v1 es:", v1)

	var v int = 30

	var p *int
	p = &v // p es un puntero que contiene la direccion de v

	fmt.Println("El puntero p apunta a la direccion:", p)
	fmt.Println("Al desferrenciar apunto al valor de v:", *p)

	fmt.Println("Antes de ejecutar la funcion v vale:", v)
	Incrementar(&v)
	fmt.Println("V ahora vale:", v)


	
}