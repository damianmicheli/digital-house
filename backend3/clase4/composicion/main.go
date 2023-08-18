package main

import "fmt"

type Autor struct {
	Nombre string
	Apellido string
}

func (a Autor) nombreCompleto() string {
	return fmt.Sprintf("%s %s", a.Nombre, a.Apellido)
}

type Libro struct {
	Titulo string
	Descripcion string
	Autor Autor
}

func (l Libro) informacion() {
	fmt.Println("Titulo:", l.Titulo)
	fmt.Println("Descripcion:", l.Descripcion)
	fmt.Println("Autor:", l.Autor.nombreCompleto())

}

func main () {
	a := Autor {"Julio", "Verne"}
	l := Libro {"Veinte mil leguas de viaje submarino", "Ciencia Ficción", a}

	l.informacion()
}
