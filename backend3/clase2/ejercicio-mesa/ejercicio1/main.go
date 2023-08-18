package main

import "fmt"

// Ejercicio 1 - Calcular salario
// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
// Categoría C: su salario es de $1.000 por hora.
// Categoría B: su salario es de $1.500 por hora, más un 20 % de su salario mensual.
// Categoría A: su salario es de $3.000 por hora, más un 50 % de su salario mensual.
// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.


func calcularSalario (minutos int, categoria string) float64 {
	
	var salario float64

	switch categoria {
	case "A":
		salario = float64(minutos / 60 * 3000 )
		salario = salario + salario * 0.5
	case "B":
		salario = float64(minutos / 60 * 1500 )
		salario = salario + salario * 0.2
	case "C":
		salario = float64(minutos / 60 * 1000 )
	}

	return salario
}

func main() {
	fmt.Println("Salario:", calcularSalario(1800, "A"))
}