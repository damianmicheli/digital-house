package main

import "fmt"

/* Registro de estudiantes
Una universidad necesita registrar a los estudiantes y generar una funcionalidad para
imprimir el detalle de los datos de cada uno de ellos, de la siguiente manera:
● Nombre: [Nombre del alumno]
● Apellido: [Apellido del alumno]
● DNI: [DNI del alumno]
● Fecha: [Fecha ingreso alumno]
Los valores que están en corchetes deben ser reemplazados por los datos brindados por los
alumnos. Para ello es necesario generar una estructura Alumno con las variables Nombre,
Apellido, DNI, Fecha y que tenga un método detalle. */


type Alumno struct {
	Nombre string
	Apellido string
	DNI string
	Fecha string
}

func (a Alumno) Detalle() {
	fmt.Printf("Nombre: %s\n", a.Nombre)
	fmt.Printf("Apellido: %s\n", a.Apellido)
	fmt.Printf("DNI: %s\n", a.DNI)
	fmt.Printf("Fecha: %s\n", a.Fecha)
}


func main () {

	alumno1 := Alumno {"Damian", "Micheli", "27123123", "30/05/2021"}

	alumno1.Detalle()

}