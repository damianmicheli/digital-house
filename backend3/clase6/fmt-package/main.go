package main

import "fmt"

func main () {
	fmt.Println(10, "Hola Mundo!")

	texto := "Mundo!"
	num := 50

	fmt.Printf("Hola %s %d\n", texto, num)

	cadena := fmt.Sprintf("Hola %s %d", texto, num)

	fmt.Println(cadena)

	var mensaje string
	fmt.Print("Ingresá un texto: ")
	fmt.Scan(&mensaje)
	fmt.Println("Mensaje:", mensaje)

	var(
		nombre string
		apellido string
		telefono int
	)
	fmt.Print("Ingresa nombre, apellido y telefono: ")
	fmt.Scanf("%s %s %d", &nombre, &apellido, &telefono)

	fmt.Printf("Nombre: %s\nApellido: %s\nTelefono: %d\n", nombre, apellido, telefono)
}
