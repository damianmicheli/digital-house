package main

// Enunciado primer parte: Impuestos de salario
// En la función main, definir una variable llamada salary y asignarle un valor de tipo “int”. Crear un error con errors.New() 
// con el mensaje “salario muy bajo" y lanzarlo en caso de que salary sea menor o igual a 10.000. La validación debe ser hecha 
// con la función Is() dentro del main.

import (
	"errors"
	"fmt"
)

var ErrLowSalary = errors.New("salario muy bajo")


func main(){

	salary := 100000

	if err := evaluateSalary(salary); err != nil {

		if errors.Is(err, ErrLowSalary) {

			fmt.Printf("el error es %v", err)

		} else {

			fmt.Printf("error desconocido")

		}
		
		return

	}
	
	fmt.Println("No hay error")

}


func evaluateSalary(salary int) error {

	const lowSalary = 10000

	if salary <= lowSalary {
		return ErrLowSalary
	}

	return nil
} 

// Repetir el proceso de la ejercitación realizada en clase, pero ahora implementando fmt.Errorf() para que el mensaje 
// de error reciba por parámetro el valor de salary indicando que no alcanza el mínimo imponible. El mensaje mostrado 
// por consola deberá decir: “Error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]” (siendo [salary] 
// 	el valor de tipo int pasado por parámetro).
