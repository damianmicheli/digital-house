package main

import "fmt"

type Compania struct {
	Nombre string
	Puesto string
}

type Empleado struct {
	Nombre string
	Apellido string
	Compania Compania
}


func (e Empleado) Informacion() {

	fmt.Printf("\nEmpleado: %s %s \nPuesto: %s\nCompañia: %s\n", e.Nombre, e.Apellido, e.Compania.Puesto, e.Compania.Nombre)

}

func (c *Compania) CambiarPuesto(nuevoPuesto string) {

	c.Puesto = nuevoPuesto

}

func main () {
	empleado := Empleado{
		Nombre: "Juan",
		Apellido: "Lopez",
		Compania: Compania{
			Nombre: "IT Sol",
			Puesto: "Jr. Dev",
		},
	}

	empleado.Informacion()
	empleado.Compania.CambiarPuesto("Sr. Dev")
	empleado.Informacion()
}