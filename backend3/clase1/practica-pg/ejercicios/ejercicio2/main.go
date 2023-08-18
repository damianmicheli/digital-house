package main

import "fmt"

func main() {

	edad := 43
	empleado := true
	antiguedad := 2
	sueldo := 230000

	resultado := validarPrestamo(edad, empleado, antiguedad, sueldo)
	fmt.Println(resultado)

	fmt.Println(suma(10,20,30))
}

func validarPrestamo (edad int, empleado bool, antiguedad int, sueldo int) string {

	var respuesta string

	if edad > 22 && empleado && antiguedad > 1 {
		if sueldo > 100000 {
			respuesta = "Prestamo aprobado sin interés."
		} else {
			respuesta = "Prestamo aprobado con interés."
				
		}
	} else {
		respuesta = "No cumple con todos los requisitos para recibir el préstamo."
	}

	return respuesta

}

func suma(valores ...float64) float64 {
	var resultado float64
	fmt.Println("el valor actual de resultado es:", resultado)
	for _,valor := range valores {
		resultado += valor
	}

	return resultado
}