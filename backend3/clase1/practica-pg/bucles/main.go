package main

import (
	"fmt"
	"time"
)

func main() {

	//FOR CLASICO
	fmt.Println("FOR CLASICO")

	sum := 0

	for i := 1; i < 100; i++ {
		sum += i
	}

	fmt.Println(sum)

	//BUCLE WHILE
	fmt.Println("\nBUCLE WHILE")

	sumWhile := 1

	for sumWhile < 10 {
		sumWhile += 1
	}

	fmt.Println(sumWhile)
	
	//bucle infinito
	fmt.Println("\nContador infinito pero hasta 10")

	contadorInfinito := 1
	fmt.Print(contadorInfinito)
	time.Sleep(time.Second)

	for contadorInfinito < 10{
		contadorInfinito++
		fmt.Print(", ",contadorInfinito)
		time.Sleep(time.Second/5)
	}

	// bucle range
	fmt.Println("\n\nbucle range")
	colores := []string{"rojo", "azul", "verde", "violeta", "naranja", "amarillo"}

	for i, color := range colores {
		fmt.Println( i, "-", color)
		time.Sleep(time.Second/5)

	}

	// continue y break se usan para saltear una iteracion o interrumpir el bucle respectivamente


}
