package main

import (
	"errors"
	"fmt"
)

// Ejercicio 2 - Calcular promedio
// Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita
// generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el
// promedio. No se pueden introducir notas negativas.

func main(){
	resultado, error := calcularPromedio(-8, 9, 7, 10)
	
	if error == nil {
		fmt.Println("Promedio:",resultado)
	} else {
		fmt.Println(error)
	}
}

func calcularPromedio(notas ...int) (float64, error) {

	cantidad := len(notas)
	sumatoria := 0

	for _, nota := range notas {
		
		if nota < 0 {
			return 0, errors.New("se ingreso una nota negativa")
		}

		sumatoria = sumatoria + nota

	}

	return float64(sumatoria) / float64(cantidad), nil
}