package main

import (
	"errors"
	"fmt"
)

func division(dividendo, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("el divisor no puede ser cero")
	}
	return dividendo/divisor, nil
}

func operaciones(valor1, valor2 float64) (float64, float64, float64, float64) {

	suma := valor1 + valor2
	resta := valor1 - valor2
	multip := valor1 * valor2
	var divis float64

	if valor2 != 0 {
		divis = valor1 / valor2
	}

	return suma, resta, multip, divis

}

func operacionesNombradas(valor1, valor2 float64) (suma float64, resta float64) {

	suma = valor1 + valor2
	resta = valor1 - valor2

	return

}

func main() {
	s, r, m, d := operaciones(6, 2)

	fmt.Println("Suma:",s)
	fmt.Println("Resta:",r)
	fmt.Println("Multiplicación:",m)
	fmt.Println("División:",d)

	sn, rn := operacionesNombradas(10,5)

	fmt.Println("Suma:",sn)
	fmt.Println("Resta:",rn)


	resultado1, error1 := division(12,6)
	fmt.Println(resultado1)
	fmt.Println(error1)

	resultado2, error2 := division(12,0)
	fmt.Println(resultado2)
	fmt.Println(error2)
}