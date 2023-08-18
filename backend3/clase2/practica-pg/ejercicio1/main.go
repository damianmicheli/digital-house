package main

import "fmt"

// Ejercicio 1 - Impuestos de salario
// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de
// depositar el sueldo.
// Para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un
// salario, teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 %
// del sueldo y si gana más de $150.000 se le descontará, además, un 10 % (27 % en total).

func calcularImpuesto(salario float64) float64 {

	var impuesto float64
	if salario > 150000 {
		impuesto = salario * 0.27
	} else if salario > 50000 {
			impuesto = salario * 0.17
	} else {
		impuesto = 0
	}

	return impuesto
}

func main() {

	salario := 200000.0
	fmt.Println("Con un salario de:", salario, "el impuesto es:", calcularImpuesto(salario))
}